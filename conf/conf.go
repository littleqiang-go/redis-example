package conf

import (
	"time"

	"github.com/BurntSushi/toml"
)

type conf struct {
	Redis *redisConf `toml:"redis"`
}

type redisConf struct {
	Addrs        []string `toml:"addrs"`
	Password     string   `toml:"password"`
	MaxRetries   int      `toml:"maxRetries"`
	PoolSize     int      `toml:"poolSize"`
	MinIdleConns int      `toml:"minIdleConns"`
	DialTimeOut  Duration `toml:"dialTimeout"`
	ReadTimeOut  Duration `toml:"readTimeout"`
	WriteTimeOut Duration `toml:"writeTimeout"`
	IdleTimeout  Duration `toml:"idleTimeout"`
	IsCluster    bool     `toml:"isCluster"`
}

type Duration time.Duration

// UnmarshalText be used toml unmarshal string time
func (d *Duration) UnmarshalText(text []byte) error {
	tp, err := time.ParseDuration(string(text))

	if err == nil {
		*d = Duration(tp)
	}

	return err
}

func RedisConf() *redisConf {
	return cfg.Redis
}

// 定义配置变量
var cfg conf

// 解析配置文件
func Parse(file string) error {
	_, err := toml.DecodeFile(file, &cfg)
	return err
}
