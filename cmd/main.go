package main

import (
	"os"

	config "github.com/alaprdfm/e-commerce/config/cfg/config"
	"github.com/alpardfm/go-toolkit/configbuilder"
	"github.com/alpardfm/go-toolkit/configreader"
	"github.com/alpardfm/go-toolkit/files"
)

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.json.template"
	appnamespace string = "e-commerce"
)

func main() {
	if !files.IsExist(configfile) {
		configbuilder.Init(configbuilder.Options{
			Env:          os.Getenv("EC_APP_ENVIRONMENT"),
			Key:          os.Getenv("EC_APP_KEY"),
			Secret:       os.Getenv("EC_APP_SECRET"),
			Region:       os.Getenv("EC_APP_REGION"),
			TemplateFile: templatefile,
			ConfigFile:   configfile,
			Namespace:    appnamespace,
		}).BuildConfig()
	}

	// init config
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		ConfigFile: configfile,
	})
	configreader.ReadConfig(&cfg)
}
