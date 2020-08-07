package dto

import "github.com/muharihar/d3ta-go/modules/auths/la/domain/schema"

// RegisterReqDTO type
type RegisterReqDTO struct {
	schema.RegisterRequest
}

// RegisterResDTO type
type RegisterResDTO struct {
	Email string `json:"email"`
}
