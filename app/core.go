package app

import (
	"github.com/BurntSushi/toml"
	"github.com/aliworkshop/echoserver"
	"github.com/aliworkshop/gateway/v2"
	"github.com/aliworkshop/logger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
)

func (a *App) initEngine() {
	a.registry.SetConfig("http.servicename", ServiceName)
	a.engine = echoserver.NewServer(a.registry)
	responder := echoserver.NewResponder(a.lang)
	controller := gateway.NewController(responder, a.mainLogger)
	a.engine.SetController(controller)
}

func (a *App) initLogger() {
	var err error
	a.mainLogger, err = logger.GetLogger(a.registry.ValueOf("logger"))
	a.panicOnErr(err)
	a.mainLogger = a.mainLogger.WithSource(ServiceName)
}

func (a *App) initLanguage() {
	type langConfig struct {
		DefaultLanguage string
		DirPath         string
		Languages       []string
	}
	conf := new(langConfig)
	a.panicOnErr(a.registry.ValueOf("language").Unmarshal(conf))
	a.lang = i18n.NewBundle(language.English)
	a.lang.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, lang := range conf.Languages {
		a.lang.MustLoadMessageFile(filepath.Join(conf.DirPath, lang))
	}
}

func (a *App) initConfig() {
	a.panicOnErr(a.registry.Unmarshal(&a.config))
	a.config.Initialize()
}
