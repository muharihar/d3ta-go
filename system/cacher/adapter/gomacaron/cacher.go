package gomacaron

import (
	"github.com/go-macaron/cache"

	_ "github.com/go-macaron/cache/ledis"
	_ "github.com/go-macaron/cache/memcache"
	_ "github.com/go-macaron/cache/mysql"
	_ "github.com/go-macaron/cache/nodb"
	_ "github.com/go-macaron/cache/postgres"
	_ "github.com/go-macaron/cache/redis"

	"github.com/muharihar/d3ta-go/system/cacher/adapter"
)

// NewCacher new GoMacaron Cacher
func NewCacher(adapter string, options cache.Options) (adapter.ICacherEngine, error) {
	var err error
	c := new(Cacher)
	c.engine, err = cache.NewCacher(adapter, options)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Cacher represent GoMacaron Cacher
type Cacher struct {
	engine cache.Cache
}

// GetEngine get default Engine
func (c *Cacher) GetEngine() interface{} {
	return c.engine
}

// IsExist check if key is exist
func (c *Cacher) IsExist(key string) bool {
	return c.engine.IsExist(key)
}

// Put put key value with timeout
func (c *Cacher) Put(key string, val interface{}, timeout int64) error {
	return c.engine.Put(key, val, timeout)
}

// Get key value
func (c *Cacher) Get(key string) interface{} {
	return c.engine.Get(key)
}

// Incr increment key value
func (c *Cacher) Incr(key string) error {
	return c.engine.Incr(key)
}

// Decr Decrement key value
func (c *Cacher) Decr(key string) error {
	return c.engine.Decr(key)
}

// Delete delete existing key value
func (c *Cacher) Delete(key string) error {
	return c.engine.Delete(key)
}

// Flush deletes all cached data
func (c *Cacher) Flush() error {
	return c.engine.Flush()
}
