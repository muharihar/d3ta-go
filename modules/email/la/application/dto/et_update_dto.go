package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
)

// ETUpdateReqDTO type
type ETUpdateReqDTO struct {
	Keys *ETUpdateKeysDTO `json:"keys"`
	Data *ETUpdateDataDTO `json:"data"`
}

// ETUpdateKeysDTO type
type ETUpdateKeysDTO struct {
	domSchema.ETUpdateKeys
}

// ETUpdateDataDTO type
type ETUpdateDataDTO struct {
	domSchema.ETUpdateData
}

// ETUpdateResDTO type
type ETUpdateResDTO struct {
	domSchema.ETUpdateResponse
}

// ToJSON covert to JSON
func (r *ETUpdateResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
