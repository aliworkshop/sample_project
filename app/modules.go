package app

import (
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/prometheus/client_golang/prometheus"
)

func (a *App) initMonitoring() {
	metric, err := a.engine.AddMonitoring(&handlerlib.Monitoring{
		Name:        "convert_methods_elapsed_time",
		Description: "time spent in convert methods",
		Type:        handlerlib.GaugeVec,
		Args:        []string{"module", "context", "name"},
	})
	metric.(*prometheus.GaugeVec).WithLabelValues("module1", "context1", "name1").Inc()
	a.panicOnErr(err)
}

func (a *App) initHelloModule() {
	a.HiModule = hello.New(a.initHr)
}
