package gomacaron

import (
	"fmt"

	"github.com/go-macaron/cache"
)

// ConfigParser config parser for gomacaron
func ConfigParser(conf interface{}) cache.Options {
	cfg := make(map[string]interface{})
	if fmt.Sprintf("%T", conf) != "map[string]interface {}" {
		cfg["adapter"] = "redis"
		cfg["adapterConfig"] = "addr:127.0.0.1:6379"
		cfg["interval"] = 60
		cfg["occupyMode"] = false
		cfg["section"] = "cache"
	} else {
		conf := conf.(map[string]interface{})
		cfg["adapter"] = conf["adapter"]
		cfg["adapterConfig"] = conf["adapterconfig"]
		cfg["interval"] = conf["interval"]
		cfg["occupyMode"] = conf["occupymode"]
		cfg["section"] = conf["section"]
	}

	return cache.Options{
		Adapter:       fmt.Sprintf("%s", cfg["adapter"]),
		AdapterConfig: fmt.Sprintf("%s", cfg["adapterConfig"]),
		Interval:      cfg["interval"].(int),
		OccupyMode:    cfg["occupyMode"].(bool),
		Section:       fmt.Sprintf("%s", cfg["section"]),
	}
}
