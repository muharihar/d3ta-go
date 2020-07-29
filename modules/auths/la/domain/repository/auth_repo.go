package repository

import (
	"github.com/muharihar/d3ta-go/modules/auths/la/domain/schema"
	"github.com/muharihar/d3ta-go/system/identity"
)

// IAuthenticationRepo interface
type IAuthenticationRepo interface {
	Register(req *schema.RegisterRequest, i identity.Identity) (*schema.RegisterResponse, error)
	ActivateRegistration(req *schema.ActivateRegistrationRequest, i identity.Identity) (*schema.ActivateRegistrationResponse, error)
	Login(req *schema.LoginRequest, i identity.Identity) (*schema.LoginResponse, error)
	LoginApp(req *schema.LoginAppRequest, i identity.Identity) (*schema.LoginAppResponse, error)
}
