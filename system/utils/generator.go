package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/blang/semver/v4"
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

// GenSemVersion generate semversion
func GenSemVersion(v string) string {
	if v == "" {
		v = "1.0.0"
	} else {
		ver, err := semver.Parse(v)
		if err != nil {
			ver, _ = semver.Parse("1.0.0")
		}
		// we use :
		//     M.m(n).p(n)
		//     1.10.10
		//     x.<+.<+
		if ver.Patch < 9 {
			ver.Patch = ver.Patch + 1
		} else {
			if ver.Minor < 9 {
				ver.Minor = ver.Minor + 1
				ver.Patch = 0
			} else {
				if ver.Major < 9 {
					ver.Major = ver.Major + 1
					ver.Minor = 0
					ver.Patch = 0
				}
			}
		}
		v = fmt.Sprintf("%v", ver)
	}

	return v
}
