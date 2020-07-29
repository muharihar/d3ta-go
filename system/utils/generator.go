package utils

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

// GenerateUUID is function to generate UUID (string)
func GenerateUUID() string {
	keyR, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return keyR.String()
}

// GeneratePassword generate password
func GeneratePassword() string {
	result := GenerateUUID() + time.Now().Format(time.RFC3339Nano)
	result = AsSha256(result)
	return result
}

// GenerateClientKey generate ClientKey
func GenerateClientKey() string {
	return GenerateUUID()
}

// GenerateSecretKey generate SecretKey
func GenerateSecretKey() string {
	return string(base64.StdEncoding.EncodeToString([]byte(AsSha256(GenerateUUID()))))
}

// GenerateRegistrationActivationCode generate RegistrationActivationCode
func GenerateRegistrationActivationCode() string {
	return GenerateUUID()
}
