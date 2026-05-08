package app

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/aliworkshop/logger"
)

func (a *App) Start() {
	a.RegisterRoutes()
	a.engine.StartMonitoring()

	go a.ChatModule.Uc.Start()

	serverErr := make(chan error, 1)
	go func() {
		a.mainLogger.InfoF("server is running on %s", a.config.Http.Address)
		serverErr <- a.engine.Run(a.config.Http.Address)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-sig:
		a.mainLogger.InfoF("received signal: %s, shutting down...", s)
	case err := <-serverErr:
		if err != nil {
			a.mainLogger.
				With(logger.Field{"error": err}).
				WithId("EngineRunError").
				ErrorF("server stopped unexpectedly")
		}
	}

	a.Stop()
}

var stopOnce sync.Once

func (a *App) Stop() {
	stopOnce.Do(func() {
		a.mainLogger.InfoF("draining chat clients...")
		a.ChatModule.Uc.Stop()

		a.mainLogger.InfoF("shutting down http server (timeout=%s)...", a.config.Http.GracefullyShutdownTimeout)
		if err := a.engine.Shutdown(a.config.Http.GracefullyShutdownTimeout); err != nil {
			a.mainLogger.
				With(logger.Field{"error": err}).
				WithId("EngineShutdownError").
				ErrorF("engine shutdown error")
		}

		a.mainLogger.InfoF("shutdown complete")
	})
}
