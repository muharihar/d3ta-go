package identity

import (
	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/system/context"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/service"
)

// Identity is subject to be refactored

// TokenType type
type TokenType string

const (
	// TokenJWT jwt Token
	TokenJWT TokenType = "JWT"
	// TokenSimple simple Token
	TokenSimple TokenType = "Simple"
)

// IdentityType Type
type IdentityType string

const (
	DefaultIdentity IdentityType = "default"
	SystemIdentity  IdentityType = "system"
)

// NewIdentity new Identity
func NewIdentity(iType IdentityType, tokenType TokenType, token string, claims *JWTCustomClaims, ctx interface{}, h *handler.Handler) (Identity, error) {

	i := Identity{
		identityType: iType,
		handler:      h,
		IsLogin:      false,
		IsAnonymous:  false,
		TokenType:    tokenType,
		Token:        token,
		Claims:       claims,
		ctx:          ctx,
	}

	jwt, err := NewJWT(h)
	if err != nil {
		return i, err
	}

	if token == "" || claims == nil {
		c, t, _, err := jwt.GenerateAnonymousToken()
		if err != nil {
			return i, err
		}
		i.Claims = c
		i.Token = t
	} else {
		i.IsLogin = true
	}

	if i.Claims.Username == AnonymousUserName {
		i.IsAnonymous = true
	}

	now := time.Now().Unix()
	if i.Claims.ExpiresAt < now {
		return i, fmt.Errorf("Token Expired")
	}
	if i.Claims.NotBefore > now {
		return i, fmt.Errorf("Token Not Valid Berofe")
	}

	i.initContextInformation()

	return i, nil
}

// Identity type
type Identity struct {
	identityType IdentityType

	ctx            interface{}
	handler        *handler.Handler
	casbinEnforcer *casbin.Enforcer

	IsLogin     bool
	IsAnonymous bool
	TokenType   TokenType
	Token       string
	Claims      *JWTCustomClaims

	ClientDevices ClientDevices
	RequestInfo   RequestInfo
}

func (i *Identity) initContextInformation() {
	// ideally we are using interface adapter
	// fmt.Printf("Type Context: `%T`\n", i.ctx)
	switch fmt.Sprintf("%T", i.ctx) {
	case "*echo.context":
		c := i.ctx.(echo.Context)

		i.ClientDevices = ClientDevices{
			UserAgent: c.Request().UserAgent(),
			IPAddress: c.RealIP(),
		}

		i.RequestInfo = RequestInfo{
			Host:          c.Request().Host,
			RemoteAddr:    c.Request().RemoteAddr,
			RequestObject: c.Request().RequestURI,
			RequestAction: c.Request().Method,
		}
	case "*context.Ctx":
		c := i.ctx.(*context.Ctx)

		i.ClientDevices = ClientDevices{
			UserAgent: c.Request().UserAgent(),
			IPAddress: c.RealIP(),
		}

		i.RequestInfo = RequestInfo{
			Host:          c.Request().HostName(),
			RemoteAddr:    c.Request().RemoteAddr(),
			RequestObject: c.Request().RequestURI(),
			RequestAction: c.Request().Method(),
		}
	default:
		i.ClientDevices = ClientDevices{
			UserAgent: "UknownCtx.UserAgent",
			IPAddress: "UknownCtx.IPAddress",
		}

		i.RequestInfo = RequestInfo{
			Host:          "UknownCtx.Host",
			RemoteAddr:    "UknownCtx.RemoteAddr",
			RequestObject: "UknownCtx.RequestObject",
			RequestAction: "UknownCtx.RequestAction",
		}
	}
}

// SetCasbinEnforcer set CasbinEnforcer
func (i *Identity) SetCasbinEnforcer(modelPath string) error {
	enf, err := service.NewCasbinEnforcer(i.handler, modelPath)
	if err != nil {
		return err
	}
	i.casbinEnforcer = enf

	return nil
}

// CanAccess can access
func (i *Identity) CanAccess(domain, obj, act string, enforcer *casbin.Enforcer) bool {
	// init enforcer
	if enforcer != nil {
		i.casbinEnforcer = enforcer
	}
	if i.casbinEnforcer == nil {
		cfg, err := i.handler.GetConfig()
		if err != nil {
			return false
		}

		// use default casbin enforcer from handler
		ceName := cfg.IAM.Casbin.Enforcers.DefaultEnforcerID
		if i.identityType == SystemIdentity {
			ceName = cfg.IAM.Casbin.Enforcers.SystemEnforcerID
		}

		ce, _ := i.handler.GetCasbinEnforcer(ceName)
		if ce == nil {
			i.casbinEnforcer, err = service.NewCasbinEnforcer(i.handler, cfg.IAM.Casbin.ModelPath)
			if err != nil {
				return false
			}
			i.handler.SetCasbinEnforcer(ceName, i.casbinEnforcer)
		} else {
			i.casbinEnforcer = ce
		}
	}

	canAccess := false
	var roles []string
	var err error

	if domain != "" {
		roles, err = i.casbinEnforcer.GetImplicitRolesForUser(i.Claims.AuthorityID, domain)
		if err != nil {
			return false
		}
	} else {
		roles, err = i.casbinEnforcer.GetImplicitRolesForUser(i.Claims.AuthorityID)
		if err != nil {
			return false
		}
	}

	for _, role := range roles {
		canAccess, err = i.casbinEnforcer.Enforce(role, obj, act)
		if err != nil {
			continue
		}
		if canAccess == true {
			break
		}
	}

	return canAccess
}

// CanAccessCurrentRequest ca access current request
func (i *Identity) CanAccessCurrentRequest() bool {
	return i.CanAccess("", i.RequestInfo.RequestObject, i.RequestInfo.RequestAction, nil)
}
