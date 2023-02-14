package monitoring

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type method struct {
	vec    *prometheus.HistogramVec
	module string
}

func newMethodHandler(subSystem, module, name, help string) *method {
	return &method{
		vec: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subSystem,
				Name:      name,
				Help:      help,
			},
			[]string{"module", "context", "name"},
		),
		module: module,
	}
}

func (h *method) MethodElapsed(ctx context.Context, name string) func() {
	ctxName := ctx.Value("name")
	start := time.Now()
	return func() {
		h.vec.WithLabelValues(h.module, fmt.Sprint(ctxName), name).Observe(time.Since(start).Seconds())
	}
}
