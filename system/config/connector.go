package config

// ConCovid19 represent Connector Covid19
type ConCovid19 struct {
	Covid19WHO  Covid19WHO  `json:"covid19who" yaml:"covid19who"`
	Covid19goid Covid19goid `json:"covid19goid" yaml:"covid19goid"`
}

// Covid19WHO represent Connector Covid19WHO - WHO
type Covid19WHO struct {
	Code   string `json:"code" yaml:"code"`
	Name   string `json:"name" yaml:"name"`
	Server string `json:"server" yaml:"server"`
	Enable bool   `json:"enable" yaml:"enable"`
}

// Covid19goid represent Connector Covid19goid - Indonesia
type Covid19goid struct {
	Code   string `json:"code" yaml:"code"`
	Name   string `json:"name" yaml:"name"`
	Server string `json:"server" yaml:"server"`
	Enable bool   `json:"enable" yaml:"enable"`
}
