package app

import (
	"os"
	"path/filepath"

	"github.com/aliworkshop/sample_project"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
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
	a.panicOnErr(a.registry.ValueOf("language").Unmarshal(&a.config.Language))
	defaultTag := language.English
	if a.config.Language.DefaultLanguage != "" {
		defaultTag = language.Make(a.config.Language.DefaultLanguage)
	}
	a.lang = i18n.NewBundle(defaultTag)
	a.lang.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	root := filepath.Join(sample_project.AppRootPath(), a.config.Language.DirPath)
	for _, lang := range a.config.Language.Languages {
		dir := filepath.Join(root, lang)
		files, err := os.ReadDir(dir)
		a.panicOnErr(err)
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			a.lang.MustLoadMessageFile(filepath.Join(dir, file.Name()))
		}
	}
}

func (a *App) initValidator() {
	a.validator = a.NewValidator().validator
}

func (a *App) initConfig() {
	a.panicOnErr(a.registry.Unmarshal(&a.config))
	a.config.Initialize()
}
