package emailtemplate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate ETUpdateRequest
func (r *ETUpdateRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Keys, validation.Required),
		validation.Field(&r.Data, validation.Required),
	)
}

// Validate ETUpdateKeys
func (r *ETUpdateKeys) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Code, validation.Length(10, 100), validation.Required),
	)
}

// Validate ETUpdateData
func (r *ETUpdateData) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.EmailFormat, validation.In("TEXT", "HTML"), validation.Required),
		validation.Field(&r.IsActive, validation.NotNil),
	)
}

// Validate ETUpdateVersion
func (r *ETUpdateVersion) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.SubjectTpl, validation.Length(10, 255), validation.Required),
		validation.Field(&r.BodyTpl, validation.Length(10, 20000), validation.Required),
	)
}
