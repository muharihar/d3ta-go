package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// ConfigParser config parser for redis
func ConfigParser(conf interface{}) redis.Options {
	cfg := make(map[string]interface{})
	if fmt.Sprintf("%T", conf) != "map[string]interface {}" {
		cfg["network"] = "tcp"
		cfg["host"] = "127.0.0.1"
		cfg["port"] = "6379"
		cfg["username"] = ""
		cfg["password"] = ""
	} else {
		conf := conf.(map[string]interface{})
		cfg["network"] = conf["network"]
		cfg["host"] = conf["host"]
		cfg["port"] = conf["port"]
		cfg["username"] = conf["username"]
		cfg["password"] = conf["password"]
		v, exist := conf["db"]
		if exist {
			cfg["db"] = v
		}
		v, exist = conf["maxretries"]
		if exist {
			cfg["maxRetries"] = v
		}
		v, exist = conf["poolsize"]
		if exist {
			cfg["poolSize"] = v
		}
		v, exist = conf["minidleconns"]
		if exist {
			cfg["minIdleConns"] = v
		}
	}

	opt := redis.Options{
		Network:  fmt.Sprintf("%s", cfg["network"]),
		Addr:     fmt.Sprintf("%s:%s", cfg["host"], cfg["port"]),
		Username: fmt.Sprintf("%s", cfg["username"]),
		Password: fmt.Sprintf("%s", cfg["password"]),
	}
	v, exist := cfg["db"]
	if exist {
		opt.DB = v.(int)
	}
	v, exist = cfg["maxRetries"]
	if exist {
		opt.MaxRetries = v.(int)
	}
	v, exist = cfg["poolSize"]
	if exist {
		opt.PoolSize = v.(int)
	}
	v, exist = cfg["minIdleConns"]
	if exist {
		opt.MinIdleConns = v.(int)
	}

	return opt
}
