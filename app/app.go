package app

import (
	od "github.com/aliworkshop/authorizer/handler/domain"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/aliworkshop/sample_project/chat"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type App struct {
	config config
	// core dependencies
	registry   configer.Registry
	engine     gateway.ServerModel
	mainLogger logger.Logger
	lang       *i18n.Bundle
	oauth      od.Handler

	HiModule   *hello.Module
	ChatModule *chat.Module
}

func New(registry configer.Registry) *App {
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
