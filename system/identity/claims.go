package identity

import (
	"github.com/dgrijalva/jwt-go"
	// "github.com/google/uuid"
)

// JWTCustomClaims represent JWT Custom claims structure
type JWTCustomClaims struct {
	ID          uint64
	UUID        string
	Username    string
	NickName    string
	AuthorityID string
	jwt.StandardClaims
}
