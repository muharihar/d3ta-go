package emailtemplate

// ETUpdateRequest type
type ETUpdateRequest struct {
	Keys *ETUpdateKeys `json:"keys"`
	Data *ETUpdateData `json:"data"`
}

// ETUpdateKeys type
type ETUpdateKeys struct {
	Code string `json:"code"`
}

// ETUpdateData type
type ETUpdateData struct {
	Name        string           `json:"name"`
	IsActive    bool             `json:"isActive"`
	EmailFormat string           `json:"emailFormat"`
	Template    *ETUpdateVersion `json:"template"`
}

// ETUpdateVersion type
type ETUpdateVersion struct {
	SubjectTpl string `json:"subjectTpl"`
	BodyTpl    string `json:"bodyTpl"`
}
