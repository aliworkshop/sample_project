package app

import (
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/aliworkshop/sample_project/monitoring"
)

func (a *App) initMonitoring() {
	m, err := monitoring.New(a.engine)
	if err != nil {
		panic(err)
	}
	a.Monitor = m
}

func (a *App) initHelloModule() {
	a.HiModule = hello.New(a.oauth)
}

func (a *App) initChatModule() {
	a.ChatModule = chat.New(a.mainLogger.With(logger.Field{
		"module": "chat",
	}))
}
