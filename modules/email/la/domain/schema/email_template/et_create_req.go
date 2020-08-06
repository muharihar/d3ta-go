package emailtemplate

// ETCreateRequest type
type ETCreateRequest struct {
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	IsActive    bool             `json:"isActive"`
	EmailFormat string           `json:"emailFormat"`
	Template    *ETCreateVersion `json:"template"`
}

// ETCreateVersion type
type ETCreateVersion struct {
	SubjectTpl string `json:"subjectTpl"`
	BodyTpl    string `json:"bodyTpl"`
}
