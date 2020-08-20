package adapter

// ICacherEngine contract/interface for CacherEngine
type ICacherEngine interface {
	GetEngine() interface{}
	IsExist(key string) bool
	Put(key string, val interface{}, timeout int64) error
	Get(key string) interface{}
	Incr(key string) error
	Decr(key string) error
	Delete(key string) error
	Flush() error
}

// CEType Cacher Engine Type
type CEType string

const (
	// CERedis Redis Cacher Engine
	CERedis CEType = "redis"
	// CEGoMacaron GoMacaron Cacher Engine
	CEGoMacaron CEType = "gomacaron"
)

// CEOptions represent Cacher Engine Options
type CEOptions struct {
	// Name of Engine adapter. Default is "CERedis (redis)".
	Engine CEType
	// Options denpent on Engine (gomacaron options, redis options)
	Options interface{}
}
