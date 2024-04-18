package main

import (
	"fmt"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/sample_project/app"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	app := app.New(config(os.Getenv("CONFIG_TYPE")))
	app.Init()
	app.InitModules()
	app.InitServices()

	app.Start()
}

func config(configType string) configer.Registry {
	if configType == "" {
		configType = "file"
	}
	var v configer.Registry
	switch strings.ToLower(configType) {
	case "remote:spring":
		configURL := os.Getenv("CONFIG_ADDRESS")
		configTOKEN := os.Getenv("CONFIG_TOKEN")
		configRepo := os.Getenv("CONFIG_REPO")
		if configURL == "" || configTOKEN == "" || configRepo == "" {
			panic("config server information not valid")
		}
		v = configer.NewSpring(configURL, configTOKEN, configRepo)
		v.SetConfigType("yaml")
		err := v.ReadConfig()
		if err != nil {
			panic(err)
		}

	case "file":
		v = configer.New()
		v.SetConfigType("yaml")
		path, _ := filepath.Abs(fmt.Sprintf("presenter/config/config.yaml"))
		f, err := os.Open(path)
		if err != nil {
			panic("cannot read config: " + err.Error())
		}
		err = v.ReadConfig(f)
		if err != nil {
			panic("cannot read config" + err.Error())
		}
	}
	return v
}
