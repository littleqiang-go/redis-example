package cache

import (
	"sync"

	"github.com/go-redis/redis"
)

var cache *redis.Client
var once sync.Once

// initialize cache
func InitCache(opts *redis.Options) error {
	once.Do(func() {
		cache = newRedisCli(opts)
	})

	_, err := cache.Ping().Result()
	return err
}

// get redis cli
func Cache() *redis.Client {
	if cache == nil {
		panic("please init redis")
	}
	return cache
}

// create redis cli
func newRedisCli(opts *redis.Options) *redis.Client {
	cli := redis.NewClient(opts)

	return cli
}

// close redis cli
func Close() error {
	if cache == nil {
		return nil
	}

	err := cache.Close()
	return err
}
