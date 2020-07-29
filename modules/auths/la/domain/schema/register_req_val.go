package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	valIs "github.com/go-ozzo/ozzo-validation/v4/is"
)

// Validate RegisterRequest
func (r *RegisterRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.Email, validation.Required, valIs.Email),
		validation.Field(&r.NickName, validation.Required),
		validation.Field(&r.Captcha, validation.Required),
		validation.Field(&r.CaptchaID, validation.Required))
}
