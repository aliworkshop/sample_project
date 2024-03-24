package app

import (
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/prometheus/client_golang/prometheus"
)

func (a *App) initMonitoring() {
	metric, err := a.engine.AddMonitoring(&gateway.Monitoring{
		Name:        "convert_methods_elapsed_time",
		Description: "time spent in convert methods",
		Type:        gateway.GaugeVec,
		Args:        []string{"module", "context", "name"},
	})
	metric.(*prometheus.GaugeVec).WithLabelValues("module1", "context1", "name1").Inc()
	a.panicOnErr(err)
}

func (a *App) initHelloModule() {
	a.HiModule = hello.New(a.oauth)
}

func (a *App) initChatModule() {
	a.ChatModule = chat.New(a.mainLogger.With(logger.Field{
		"module": "chat",
	}))
}
