package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/muharihar/d3ta-go/system/cacher/adapter"
)

// NewCacher new Redis Cacher
func NewCacher(options *redis.Options) (adapter.ICacherEngine, error) {
	c := new(Cacher)
	c.ctx = context.Background()
	c.engine = redis.NewClient(options)
	return c, nil
}

// Cacher represent Redis Cacher
type Cacher struct {
	engine *redis.Client
	ctx    context.Context
}

// GetEngine get default Engine
func (c *Cacher) GetEngine() interface{} {
	return c.engine
}

// IsExist check key is exist
func (c *Cacher) IsExist(key string) bool {
	r, err := c.engine.Exists(c.ctx, key).Result()
	if err != nil {
		return false
	}
	return r > 0
}

// Put put key value with (timeout) expiration
func (c *Cacher) Put(key string, val interface{}, timeout int64) error {
	expiration, err := time.ParseDuration(fmt.Sprintf("%ds", timeout))
	if err != nil {
		return err
	}
	return c.engine.Set(c.ctx, key, val, expiration).Err()
}

// Get key value
func (c *Cacher) Get(key string) interface{} {
	r, err := c.engine.Get(c.ctx, key).Result()
	if err != nil {
		return nil
	}
	return r
}

// Incr increment key value
func (c *Cacher) Incr(key string) error {
	return c.engine.Incr(c.ctx, key).Err()
}

// Decr decrement key value
func (c *Cacher) Decr(key string) error {
	return c.engine.Decr(c.ctx, key).Err()
}

// Delete delete key
func (c *Cacher) Delete(key string) error {
	return c.engine.Del(c.ctx, key).Err()
}

// Flush flush all data
func (c *Cacher) Flush() error {
	return c.engine.FlushAll(c.ctx).Err()
}
