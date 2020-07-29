package adapter

// Covid19AdapterInfo represent Covid19AdapterInfo
type Covid19AdapterInfo struct {
	Code   string `yaml:"code" json:"code"`
	Name   string `yaml:"name" json:"name"`
	Server string `yaml:"server" json:"server"`
	Enable bool   `yaml:"enable" json:"enable"`
}
