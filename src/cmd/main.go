package main

import (
	"os"

	"github.com/alpardfm/e-commerce/src/business/domain"
	"github.com/alpardfm/e-commerce/src/business/usecase"
	"github.com/alpardfm/e-commerce/src/handler/rest"
	"github.com/alpardfm/e-commerce/src/utils/config"
	"github.com/alpardfm/go-toolkit/configbuilder"
	"github.com/alpardfm/go-toolkit/configreader"
	"github.com/alpardfm/go-toolkit/files"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/parser"
	"github.com/alpardfm/go-toolkit/sql"
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

	// init logger
	log := log.Init(cfg.Log)

	// init parser
	parser := parser.InitParser(log, cfg.Parser)

	//json paster
	JSONParser := parser.JSONParser()

	// init db conn
	db := sql.Init(cfg.SQL, log)

	// init all domain
	d := domain.Init(log, db, JSONParser, cfg)

	// init all uc
	uc := usecase.Init(log, d, JSONParser, cfg)

	// init and run http server
	r := rest.Init(cfg, configreader, log, parser.JSONParser(), uc)
	r.Run()

}
