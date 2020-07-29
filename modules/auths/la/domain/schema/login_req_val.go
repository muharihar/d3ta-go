package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate LoginRequest
func (r *LoginRequest) Validate() error {

	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.Captcha, validation.Required),
		validation.Field(&r.CaptchaID, validation.Required),
	)
}
