package main

import (
	"flag"
	"os"

	"com.github.sunbenxin/config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// global var
var (
	conf *config.ServerConfig = &config.ServerConfig{}
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "../config/server.yaml", "Config file")
}

func main() {
	log.Info("start backend...")
	defer func() {
		if err := recover(); err != nil {
			log.Infof("panic recover err:%v", err)
		}
	}()

	initConfig()
	// init DB
	// init server
	// start server
}

func initConfig() {
	log.Infof("config file: %v", configFile)
	if !flag.Parsed() {
		flag.Parse()
	}

	cfgBytes, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("read file %v failed,err %v", configFile, err)
	}
	if err := yaml.Unmarshal(cfgBytes, conf); err != nil {
		log.Fatalf("yaml parse config file failed,err %v", err)
	}
	log.Infof("server config %+v, db config %+v", conf, conf.DB)
}
