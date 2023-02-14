package app

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/echoserver"
	"github.com/aliworkshop/handlerlib"
	"github.com/aliworkshop/loggerlib/logger"
	od "github.com/aliworkshop/oauthlib/handler/domain"
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

	HiModule *hello.Module
}

func New(registry configlib.Registry) *App {
	return &App{registry: registry}
}

func (a *App) Init() {
	a.initConfig()
	a.initLogger()
	a.initEngine()
	a.initLanguage()
}

func (a *App) InitModules() {
	a.initHelloModule()
	a.initMonitoring()
}

func (a *App) InitServices() {
	a.initOauth()
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
