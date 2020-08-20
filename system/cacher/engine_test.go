package cacher

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/cacher/adapter"
	ceGoMacaron "github.com/muharihar/d3ta-go/system/cacher/adapter/gomacaron"
	ceRedis "github.com/muharihar/d3ta-go/system/cacher/adapter/redis"
)

func TestEngineRedis_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
	}

	e, ct, err := NewCacherEngine(RedisCacher, adapter.CEOptions{
		Engine: adapter.CERedis,
		/*
			Options: redis.Options{
				Network: "tcp",
				Addr:    "127.0.0.1:6379",
				...
			},
		*/
		Options: ceRedis.ConfigParser(cfg.Caches.SessionCache.Configurations),
	})

	if err != nil {
		t.Errorf("NewCacherEngine: %v", err.Error())
	}

	c, err := NewCacher(ct, e)
	if err != nil {
		t.Errorf("NewCacher: %v", err.Error())
	}

	c.Context = "Context"
	c.Container = "Container"
	c.Component = "Component"

	testCacher(c, t)
}

func TestEngineGoMacaronRedis_Methods(t *testing.T) {
	cfg, err := newConfig(t)
	if err != nil {
		t.Errorf("newConfig: %v", err.Error())
	}

	ce, ct, err := NewCacherEngine(RedisCacher, adapter.CEOptions{
		Engine: adapter.CEGoMacaron,
		/*
			Options: cache.Options{
				Adapter:       "redis",
				AdapterConfig: "addr:127.0.0.1:6379",
				...
			},
		*/
		Options: ceGoMacaron.ConfigParser(cfg.Caches.TmpDataCache.Configurations),
	})

	if err != nil {
		t.Errorf("NewCacherEngine: %v", err.Error())
	}

	c, err := NewCacher(ct, ce)
	if err != nil {
		t.Errorf("NewCacher: %v", err.Error())
	}

	c.Context = "Context"
	c.Container = "Container"
	c.Component = "Component"

	testCacher(c, t)
}
