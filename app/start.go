package app

import (
	"github.com/aliworkshop/loggerlib/logger"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	a.RegisterRoutes()
	a.engine.StartMonitoring()
	go func() {
		a.ChatModule.Uc.Start()
	}()
	go func() {
		a.mainLogger.InfoF("server is running on :%s", a.config.Http.Address)
		if err := a.engine.Run(a.config.Http.Address); err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	a.Stop()
}

func (a *App) Stop() {
	a.mainLogger.InfoF("shutting down...")
	err := a.engine.Shutdown(a.config.Http.GracefullyShutdownTimeout)
	if err != nil {
		a.mainLogger.
			With(logger.Field{
				"error": err,
			}).
			WithId("EngineShutdownError").
			ErrorF("engine shutdown error")
	}
	a.mainLogger.InfoF("shutdown")
}
