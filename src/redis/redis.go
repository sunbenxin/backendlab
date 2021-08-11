// Package redis provides ...
package redis

import (
	"com.github.sunbenxin/config"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client

var rlock *redsync.Redsync

// Init ...
func Init(cfg *config.RedisConfig) {
	rdb = redis.NewClient(&redis.Options{
		Addr: cfg.Host,
	})

	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	// 多个pool支持quorum
	pool := goredis.NewPool(rdb)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rlock = redsync.New(pool)
	log.Info("init redis success")

}

// test func for rlock
func test() {
	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutexname := "my-global-mutex"
	mutex := rlock.NewMutex(mutexname)

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}
