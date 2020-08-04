package emailtemplate

// ETSetActiveRequest type
type ETSetActiveRequest struct {
	Keys *ETSetActiveKeys `json:"keys"`
	Data *ETSetActiveData `json:"data"`
}

// ETSetActiveKeys type
type ETSetActiveKeys struct {
	Code string `json:"code"`
}

// ETSetActiveData type
type ETSetActiveData struct {
	IsActive bool `json:"isActive"`
}
