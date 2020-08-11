package config

// Connectors represent available connectors
type Connectors struct {
	Identity Identity   `json:"identity" yaml:"identity"`
	Covid19  ConCovid19 `json:"covid19" yaml:"covid19"`
}

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

// Identity type
type Identity struct {
	EA2M EA2M `json:"ea2m" yaml:"ea2m"`
}

// EA2M Type
type EA2M struct {
	Server            string `json:"server" yaml:"server"`
	ClientAccessKey   string `json:"clientAccessKey" yaml:"clientAccessKey"`
	ClientSecretKey   string `json:"clientSecretKey" yaml:"clientSecretKey"`
	AllowDevToken     bool   `json:"allowDevToken" yaml:"allowDevToken"`
	DevIdentityToken  string `json:"devIdentityToken" yaml:"devIdentityToken"`
	DevIdentityClaims string `json:"devIdentityClaims" yaml:"devIdentityClaims"`
}
