package cacher

import (
	"fmt"

	"github.com/muharihar/d3ta-go/system/cacher/adapter"
)

// CacherType Represent CacherType
type CacherType string

const (
	MemoryCacher   CacherType = "memory"
	FileCacher     CacherType = "file"
	RedisCacher    CacherType = "redis"
	MemcacheCacher CacherType = "memcache"
	PGCacher       CacherType = "postgres"
	MyCacher       CacherType = "mysql"
	LedisCacher    CacherType = "ledis"
	NoDBCacher     CacherType = "nodb"
)

// NewCacher create new Cacher
func NewCacher(cacherType CacherType, cacherEngine adapter.ICacherEngine) (*Cacher, error) {

	if cacherEngine == nil {
		return nil, fmt.Errorf("Invalid cacherEngine value")
	}

	c := Cacher{
		_type:        cacherType,
		cacherEngine: cacherEngine,
	}

	// C4 prefix key
	c.Context = "defaultContext"
	c.Container = "defaultContainer"
	c.Component = "defaultComponent"

	return &c, nil
}

// Cacher type
type Cacher struct {
	_type        CacherType
	cacherEngine adapter.ICacherEngine

	Context   string
	Container string
	Component string
}

// GetEngine get default Engine
func (c *Cacher) GetEngine() interface{} {
	return c.cacherEngine.GetEngine()
}

// IsExist check if key is exist
func (c *Cacher) IsExist(key string) bool {
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.IsExist(_key)
}

// Put put key value with timeout
func (c *Cacher) Put(key string, val interface{}, timeout int64) error {
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.Put(_key, val, timeout)
}

// Get key value
func (c *Cacher) Get(key string) interface{} {
	if !c.IsExist(key) {
		return nil
	}
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.Get(_key)
}

// Incr increment key value
func (c *Cacher) Incr(key string) error {
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.Incr(_key)
}

// Decr Decrement key value
func (c *Cacher) Decr(key string) error {
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.Decr(_key)
}

// Delete delete existing key value
func (c *Cacher) Delete(key string) error {
	if !c.IsExist(key) {
		return nil
	}
	_key := fmt.Sprintf("%s.%s.%s.%s", c.Context, c.Container, c.Component, key)
	return c.cacherEngine.Delete(_key)
}

// Flush deletes all cached data
func (c *Cacher) Flush() error {
	return c.cacherEngine.Flush()
}
