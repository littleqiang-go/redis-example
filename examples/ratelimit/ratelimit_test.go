package ratelimit

import (
	"redis-example/cache"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func Test_LimitIP(t *testing.T) {
	cache.InitCache(&redis.Options{Addr: "localhost:6379"})
	const ip = "127.0.0.1"
	for i := 1; i <= 50; i++ {
		result, err := WindowRateLimit(ip, 1, 10)
		if err != nil {
			t.Error(err)
			return
		}
		if i == 11 {
			if !result {
				t.Logf("%s ip limit => %v", ip, result)
			} else {
				t.Errorf("%s cross %v", ip, result)
			}
		} else {
			t.Logf("%s ip result => %v", ip, result)
		}
		time.Sleep(time.Millisecond * 50)
	}

}
