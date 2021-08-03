package main

import (
	"flag"
	"log"
	cache "redis-example/cache"
	"redis-example/conf"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	Init()
}

// init
func Init() {
	var cfgFile string
	flag.StringVar(&cfgFile, "c", "config.toml", "config path")

	err := conf.Parse(cfgFile)
	if err != nil {
		panic(err)
	}
	initCache()
}

// initialize cache
func initCache() {
	err := cache.InitCache(&redis.Options{
		Addr:         conf.RedisConf().Addrs[0],
		Password:     conf.RedisConf().Password,
		MaxRetries:   conf.RedisConf().MaxRetries,
		PoolSize:     conf.RedisConf().PoolSize,
		MinIdleConns: conf.RedisConf().MinIdleConns,
		DialTimeout:  time.Duration(conf.RedisConf().DialTimeOut),
		ReadTimeout:  time.Duration(conf.RedisConf().ReadTimeOut),
		WriteTimeout: time.Duration(conf.RedisConf().WriteTimeOut),
		IdleTimeout:  time.Duration(conf.RedisConf().IdleTimeout),
	})

	if err != nil {
		panic(err)
	}

	log.Println("redis init suc")
}
