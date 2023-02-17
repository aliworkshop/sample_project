package app

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/echoserver"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/loggerlib/logger"
	od "github.com/aliworkshop/oauthlib/handler/domain"
	"github.com/aliworkshop/sample_project/chat"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type App struct {
	config config
	// core dependencies
	registry   configlib.Registry
	engine     handlerlib.ServerModel
	mainLogger logger.Logger
	lang       *i18n.Bundle
	oauth      od.Handler

	HiModule   *hello.Module
	ChatModule *chat.Module
}

func New(registry configlib.Registry) *App {
	return &App{registry: registry}
}

func (a *App) Init() {
	a.initConfig()
	a.initLogger()
	a.initEngine()
	a.initLanguage()
	a.initOauth()
}

func (a *App) InitModules() {
	a.initHelloModule()
	a.initChatModule()
	a.initMonitoring()
}

func (a *App) InitServices() {
}

func (a *App) panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (a *App) initHr() handlerlib.HandlerModel {
	responder := echoserver.NewResponder(a.lang)
	baseHandler := echoserver.NewHandler(a.mainLogger)
	return handlerlib.NewModel(baseHandler, responder)
}
