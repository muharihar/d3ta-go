package cacher

import (
	"testing"

	"github.com/go-macaron/cache"
	"github.com/go-redis/redis/v8"
	"github.com/muharihar/d3ta-go/system/cacher/adapter"
)

func TestEngineRedis_Methods(t *testing.T) {
	e, ct, err := NewCacherEngine(RedisCacher, adapter.CEOptions{
		Engine: adapter.CERedis,
		Options: redis.Options{
			Network: "tcp",
			Addr:    "127.0.0.1:6379",
		},
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
	ce, ct, err := NewCacherEngine(RedisCacher, adapter.CEOptions{
		Engine: adapter.CEGoMacaron,
		Options: cache.Options{
			Adapter:       "redis",
			AdapterConfig: "addr:127.0.0.1:6379",
		},
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
