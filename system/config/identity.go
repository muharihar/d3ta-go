package config

// IAM type
type IAM struct {
	DefaultAdmin DefaultAdmin `json:"defaultAdmin" yaml:"defaultAdmin"`
	Registration Registration `json:"registration" yaml:"registration"`
	JWT          JWT          `json:"JWT" yaml:"JWT"`
	Casbin       Casbin       `json:"casbin" yaml:"casbin"`
}

// DefaultAdmin type
type DefaultAdmin struct {
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	NickName    string `json:"nickName" yaml:"nickName"`
	Email       string `json:"email" yaml:"email"`
	AuthorityID string `json:"authorityID" yaml:"authorityID"`
}

// Registration type
type Registration struct {
	ActivationURL      string `json:"activationURL" yaml:"activationURL"`
	DefaultAuthorityID string `json:"defaultAuthorityID" yaml:"defaultAuthorityID"`
}

// JWT Type
type JWT struct {
	Issuer     string `json:"issuer" yaml:"issuer"`
	SigningKey string `json:"signingKey" yaml:"signingKey"`
}

// Casbin Type
type Casbin struct {
	ModelPath string          `json:"modelPath" yaml:"modelPath"`
	Enforcers CasbinEnforcers `json:"enforcers" yaml:"enforcers"`
}

// CasbinEnforcers type
type CasbinEnforcers struct {
	DefaultEnforcerID string `json:"defaultEnforcerID" yaml:"defaultEnforcerID"`
	SystemEnforcerID  string `json:"systemEnforcerID" yaml:"systemEnforcerID"`
}
