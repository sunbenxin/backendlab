// Package database provides ...
package database

import (
	"fmt"

	"com.github.sunbenxin/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDb ...
func InitDB(dbCfg *config.DbConfig) {
	// dsn format: username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/backend", dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port)
	log.Infof("init db dsn: %s", dsn)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("init db err:%+v", err)
	}
	log.Info("init db success...")
}
