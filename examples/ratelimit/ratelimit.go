package ratelimit

import (
	"redis-example/cache"
	"strconv"
	"time"
)

// 滑动窗口限流
// limitKey 可以使接口名，ip等
// limitTime 限流时间 单位秒
// windowSize 窗口大小
func WindowRateLimit(limitKey string, limitTime, windowSize int) (bool, error) {
	rdb := cache.Cache()
	now := time.Now().Unix()
	length, err := rdb.LLen(limitKey).Result()

	if err != nil {
		return false, err
	}

	if length <= int64(windowSize) {
		_, err = rdb.RPush(limitKey, now).Result()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	first, err := rdb.LIndex(limitKey, 0).Result()
	if err != nil {
		return false, err
	}

	firstTime, err := strconv.ParseInt(first, 10, 64)
	if err != nil {
		return false, err
	}

	if now-firstTime <= int64(limitTime) {
		return false, nil
	} else {
		_, err = rdb.LPop(limitKey).Result()
		if err != nil {
			return false, err
		}
		_, err = rdb.RPush(limitKey, now).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
