package cacher

import (
	"fmt"

	"github.com/go-macaron/cache"
	"github.com/go-redis/redis/v8"
	"github.com/muharihar/d3ta-go/system/cacher/adapter"
	ceGoMacaron "github.com/muharihar/d3ta-go/system/cacher/adapter/gomacaron"
	ceRedis "github.com/muharihar/d3ta-go/system/cacher/adapter/redis"
)

// NewCacherEngine new Cacher Engine
func NewCacherEngine(cacherType CacherType, ceOptions adapter.CEOptions) (adapter.ICacherEngine, CacherType, error) {
	tOpt := fmt.Sprintf("%T", ceOptions.Options)

	switch ceOptions.Engine {
	case adapter.CERedis:
		if tOpt != "redis.Options" {
			return nil, "", fmt.Errorf("Invalid Redis options (should be `redis.Options`)")
		}
		opt := ceOptions.Options.(redis.Options)
		ce, err := ceRedis.NewCacher(&opt)
		if err != nil {
			return nil, "", err
		}
		return ce, cacherType, nil

	case adapter.CEGoMacaron:
		if tOpt != "cache.Options" {
			return nil, "", fmt.Errorf("Invalid GoMacaron cache options (should be `cache.Options`)")
		}
		opt := ceOptions.Options.(cache.Options)
		ce, err := ceGoMacaron.NewCacher(string(cacherType), opt)
		if err != nil {
			return nil, "", err
		}
		return ce, cacherType, nil

	default:
		return nil, "", fmt.Errorf("Invalid Cacher Engine")
	}
}
