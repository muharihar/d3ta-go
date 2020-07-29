package config

// Securities type
type Securities struct {
	Captcha Captcha `json:"captcha" yaml:"captcha"`
}

// Captcha type
type Captcha struct {
	KeyLong   int    `json:"keyLong" yaml:"keyLong"`
	ImgWidth  int    `json:"imgWidth" yaml:"imgWidth"`
	ImgHeight int    `json:"imgHeight" yaml:"imgHeight"`
	ImgURL    string `json:"imgURL" yaml:"imgURL"`
}
