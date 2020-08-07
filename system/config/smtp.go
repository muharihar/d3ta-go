package config

// SMTPServers type
type SMTPServers struct {
	DefaultSMTP SMTP `json:"defaultSMTP" yaml:"defaultSMTP"`
}

// SMTP Type
type SMTP struct {
	Server      string `json:"server" yaml:"server"`
	Port        string `json:"port" yaml:"port"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	SenderEmail string `json:"senderEmail" yaml:"senderEmail"`
	SenderName  string `json:"senderName" yaml:"senderName"`
}
