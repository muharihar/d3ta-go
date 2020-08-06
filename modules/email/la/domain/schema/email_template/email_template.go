package emailtemplate

// EmailTemplate type
type EmailTemplate struct {
	ID               uint64 `json:"ID"`
	UUID             string `json:"uuid"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	IsActive         bool   `json:"isActive"`
	EmailFormat      string `json:"emailFormat"`
	DefaultVersionID uint64 `json:"defaultVersionID"`
}

// EmailTemplateVersion type
type EmailTemplateVersion struct {
	ID              uint64         `json:"ID"`
	Version         string         `json:"version"`
	SubjectTpl      string         `json:"subjectTpl"`
	BodyTpl         string         `json:"bodyTpl"`
	EmailTemplateID uint64         `json:"emailTemplateID"`
	EmailTemplate   *EmailTemplate `json:"emailTemplate"`
}
