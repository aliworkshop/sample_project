package tests

import (
	"fmt"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/sample_project"
	"github.com/aliworkshop/sample_project/app"
	"os"
	"path/filepath"
)

func testApp() *app.App {
	app := app.New(testConfig())
	app.Init()
	app.InitModules()
	app.InitServices()
	return app
}

func testConfig() configer.Registry {
	v := configer.New()
	v.SetConfigType("yaml")
	path, _ := filepath.Abs(fmt.Sprintf("%s/presenter/config/config-test.yaml", sample_project.AppRootPath()))
	f, err := os.Open(path)
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = v.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}
	return v
}
