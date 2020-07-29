package covid19goid

// Config represent Config
type Config struct {
	Code   string `yaml:"code" json:"code"`
	Name   string `yaml:"name" json:"name"`
	Server string `yaml:"server" json:"server"`
	Enable bool   `yaml:"enable" json:"enable"`
}
