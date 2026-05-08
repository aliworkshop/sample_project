package monitoring

import (
	"context"
	"strconv"
	"time"

	"github.com/aliworkshop/errors"
	"github.com/aliworkshop/gateway/v2"
	"github.com/prometheus/client_golang/prometheus"
)

const subsystem = "app"

type Monitor struct {
	requestsTotal    *prometheus.CounterVec
	requestDuration  *prometheus.HistogramVec
	requestsInFlight *prometheus.GaugeVec
	methodDuration   *prometheus.HistogramVec
}

func New(engine gateway.ServerModel) (*Monitor, errors.ErrorModel) {
	m := &Monitor{}

	requestsTotal, err := engine.AddMonitoring(&gateway.Monitoring{
		Subsystem:   subsystem,
		Name:        "http_requests_total",
		Description: "total HTTP requests handled, partitioned by module, handler and status code",
		Type:        gateway.CounterVec,
		Args:        []string{"module", "handler", "status"},
	})
	if err != nil {
		return nil, err
	}
	m.requestsTotal = requestsTotal.(*prometheus.CounterVec)

	requestDuration, err := engine.AddMonitoring(&gateway.Monitoring{
		Subsystem:   subsystem,
		Name:        "http_request_duration_seconds",
		Description: "HTTP request handler duration in seconds",
		Type:        gateway.HistogramVec,
		Args:        []string{"module", "handler"},
		Buckets:     prometheus.DefBuckets,
	})
	if err != nil {
		return nil, err
	}
	m.requestDuration = requestDuration.(*prometheus.HistogramVec)

	inFlight, err := engine.AddMonitoring(&gateway.Monitoring{
		Subsystem:   subsystem,
		Name:        "http_requests_in_flight",
		Description: "HTTP requests currently being handled, partitioned by module and handler",
		Type:        gateway.GaugeVec,
		Args:        []string{"module", "handler"},
	})
	if err != nil {
		return nil, err
	}
	m.requestsInFlight = inFlight.(*prometheus.GaugeVec)

	methodDuration, err := engine.AddMonitoring(&gateway.Monitoring{
		Subsystem:   subsystem,
		Name:        "method_duration_seconds",
		Description: "duration of an instrumented method in seconds",
		Type:        gateway.HistogramVec,
		Args:        []string{"module", "context", "name"},
		Buckets:     prometheus.DefBuckets,
	})
	if err != nil {
		return nil, err
	}
	m.methodDuration = methodDuration.(*prometheus.HistogramVec)

	return m, nil
}

// Wrap returns a gateway.Handler that records HTTP-level metrics around the
// underlying handler. The (module, handler) labels become the dashboard
// dimensions in Grafana.
func (m *Monitor) Wrap(module, handler string, h gateway.Handler) gateway.Handler {
	if m == nil {
		return h
	}
	return &instrumented{mon: m, module: module, name: handler, inner: h}
}

// MethodElapsed times an arbitrary block of code. Idiomatic use:
//
//	defer mon.MethodElapsed(ctx, "FetchUser")()
func (m *Monitor) MethodElapsed(ctx context.Context, name string) func() {
	if m == nil {
		return func() {}
	}
	module, _ := ctx.Value(ctxKeyModule).(string)
	contextName, _ := ctx.Value(ctxKeyContext).(string)
	start := time.Now()
	return func() {
		m.methodDuration.WithLabelValues(module, contextName, name).Observe(time.Since(start).Seconds())
	}
}

type ctxKey string

const (
	ctxKeyModule  ctxKey = "monitoring.module"
	ctxKeyContext ctxKey = "monitoring.context"
)

// WithLabels returns a context carrying the module/context labels picked up by
// MethodElapsed.
func WithLabels(ctx context.Context, module, contextName string) context.Context {
	ctx = context.WithValue(ctx, ctxKeyModule, module)
	ctx = context.WithValue(ctx, ctxKeyContext, contextName)
	return ctx
}

type instrumented struct {
	mon    *Monitor
	module string
	name   string
	inner  gateway.Handler
}

func (i *instrumented) Handle(req gateway.HttpRequester) (any, errors.ErrorModel) {
	i.mon.requestsInFlight.WithLabelValues(i.module, i.name).Inc()
	defer i.mon.requestsInFlight.WithLabelValues(i.module, i.name).Dec()

	start := time.Now()
	res, err := i.inner.Handle(req)
	elapsed := time.Since(start).Seconds()

	status := strconv.Itoa(req.GetStatusCode())
	if err != nil {
		// gateway sets the response status during Respond, after the handler
		// returns; fall back to a stable label so failures still group.
		status = "error"
	}
	i.mon.requestsTotal.WithLabelValues(i.module, i.name, status).Inc()
	i.mon.requestDuration.WithLabelValues(i.module, i.name).Observe(elapsed)
	return res, err
}
