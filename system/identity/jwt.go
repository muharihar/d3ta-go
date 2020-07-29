package identity

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/muharihar/d3ta-go/system/handler"
)

// JWT type
type JWT struct {
	handler    *handler.Handler
	SigningKey []byte
	Issuer     string
}

var (
	ErrTokenExpired     = errors.New("Token is expired [TokenExpired]")
	ErrTokenNotValidYet = errors.New("Token not active yet [TokenNotValidYet]")
	ErrTokenMalformed   = errors.New("That's not even a token [TokenMalformed]")
	ErrTokenInvalid     = errors.New("Couldn't handle this token [IvalidToken]")
)

// NewJWT new JWT
func NewJWT(h *handler.Handler) (*JWT, error) {
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}

	return &JWT{
		handler:    h,
		SigningKey: []byte(cfg.IAM.JWT.SigningKey),
		Issuer:     cfg.IAM.JWT.Issuer}, nil
}

// CreateToken Create a token
func (j *JWT) CreateToken(claims JWTCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken Parse token
func (j *JWT) ParseToken(tokenString string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid

	}

}

// RefreshToken update token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}

// GenerateAnonymousToken generate anonymous token
func (j *JWT) GenerateAnonymousToken() (c *JWTCustomClaims, token string, expiredAt int64, err error) {

	claims := j.CreateCustomClaims(AnonymousID, AnonymousUUID, AnonymousUserName, AnonymousNickName, AnonymousAuthorityID)

	token, expiredAt, err = j.GenerateToken(claims)

	return &claims, token, expiredAt, err
}

// GenerateToken generate token
func (j *JWT) GenerateToken(claims JWTCustomClaims) (token string, expiredAt int64, err error) {

	token, err = j.CreateToken(claims)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to generate token: %v", err.Error())
	}
	expiredAt = claims.StandardClaims.ExpiresAt * 1000

	return token, expiredAt, nil
}

// CreateCustomClaims create custom claims
func (j *JWT) CreateCustomClaims(ID uint64, UUID, username, nickName, authorityID string) JWTCustomClaims {

	return JWTCustomClaims{
		ID:          ID,
		UUID:        UUID,
		Username:    username,
		NickName:    nickName,
		AuthorityID: authorityID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // signature effective time
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // expiration time one week
			Issuer:    j.Issuer,                       // Issuer of the signature
		},
	}
}
