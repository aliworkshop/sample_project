package app

import (
	od "github.com/aliworkshop/authorizer/handler/domain"
	"github.com/aliworkshop/sample_project/chat"
	"github.com/aliworkshop/sample_project/hello"
	"github.com/aliworkshop/sample_project/monitoring"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type App struct {
	config config
	// core dependencies
	registry   configer.Registry
	engine     gateway.ServerModel
	mainLogger logger.Logger
	lang       *i18n.Bundle
	validator  *validator.Validate
	oauth      od.Handler

	HiModule   *hello.Module
	ChatModule *chat.Module

	Monitor *monitoring.Monitor
}

func New(registry configer.Registry) *App {
	return &App{registry: registry}
}

func (a *App) Init() {
	a.initConfig()
	a.initLogger()
	a.initLanguage()
	a.initEngine()
	a.initValidator()
	a.initOauth()
}

func (a *App) InitModules() {
	a.initMonitoring()
	a.initHelloModule()
	a.initChatModule()
}

func (a *App) InitServices() {
}

func (a *App) panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
