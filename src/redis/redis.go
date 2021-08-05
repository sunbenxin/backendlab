// Package redis provides ...
package redis

import (
	"com.github.sunbenxin/config"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

// Init ...
func Init(cfg *config.RedisConfig) {
	rdb = redis.NewClient(&redis.Options{
		Addr: cfg.Host,
	})
	log.Info("init redis success")
}
