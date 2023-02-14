package main

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/sample_project/app"
	"os"
	"strings"
)

func main() {
	app := app.New(config(os.Getenv("CONFIG_TYPE")))
	app.Init()
	app.InitModules()
	app.InitServices()

	app.Start()
}

func config(configType string) configlib.Registry {
	var v configlib.Registry
	switch strings.ToLower(configType) {
	case "remote:spring":
		configURL := os.Getenv("CONFIG_ADDRESS")
		configTOKEN := os.Getenv("CONFIG_TOKEN")
		configRepo := os.Getenv("CONFIG_REPO")
		if configURL == "" || configTOKEN == "" || configRepo == "" {
			panic("config server information not valid")
		}
		v = configlib.NewSpring(configURL, configTOKEN, configRepo)
		v.SetConfigType("yaml")
		err := v.ReadConfig()
		if err != nil {
			panic(err)
		}

	case "file":
		v = configlib.New()
		v.SetConfigType("yaml")
		f, err := os.Open("./config.yaml")
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
