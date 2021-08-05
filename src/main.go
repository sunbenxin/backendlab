package main

import (
	"flag"
	"fmt"
	"os"

	"com.github.sunbenxin/config"
	ctl "com.github.sunbenxin/controller"
	"com.github.sunbenxin/database"
	"com.github.sunbenxin/redis"
	"gopkg.in/yaml.v3"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	database.InitDB(conf.DB)
	redis.Init(conf.Redis)
	startServer()
}

func startServer() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", ctl.Hello)
	}

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	log.Infof("server addr:%s", addr)
	r.Run(addr)
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
