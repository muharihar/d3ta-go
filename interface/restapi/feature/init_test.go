package feature

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fsnotify/fsnotify"
	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

func newConfig(h *handler.Handler) (*config.Config, error) {
	configPath := "../../../conf"

	//init config
	cfg, viper, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}
	cfg.IAM.Casbin.ModelPath = "../../../conf/casbin/casbin_rbac_rest_model.conf"

	h.SetConfig(cfg)

	viper.OnConfigChange(func(e fsnotify.Event) {
		c := new(config.Config)
		if err := viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}
		h.SetConfig(c)
	})

	return cfg, nil
}

func newHandler() *handler.Handler {

	h, _ := handler.NewHandler()

	// init configuration
	_, err := newConfig(h)
	if err != nil {
		panic(err)
	}

	return h
}

func generateUserTestToken(h *handler.Handler, t *testing.T) (string, *identity.JWTCustomClaims, error) {
	j, err := identity.NewJWT(h)
	if err != nil {
		return "", nil, err
	}

	claims := identity.JWTCustomClaims{
		ID:          0,
		UUID:        "test-test-test-test-test",
		Username:    "test.d3tago",
		NickName:    "Test User",
		AuthorityID: "group:admin",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,           // signature effective time
			ExpiresAt: time.Now().Unix() + 60*60*24*30*12, // expiration time 12 month
			Issuer:    j.Issuer,
		},
	}

	token, _, err := j.GenerateToken(claims)
	if err != nil {
		return "", nil, err
	}

	return token, &claims, nil
}
