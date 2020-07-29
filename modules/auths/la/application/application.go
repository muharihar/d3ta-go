package application

import (
	"github.com/muharihar/d3ta-go/modules/auths/la/application/service"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewAuthsApp new AuthsApp
func NewAuthsApp(h *handler.Handler) (*AuthsApp, error) {
	var err error

	app := new(AuthsApp)
	app.handler = h

	if app.AuthenticationSvc, err = service.NewAuthenticationSvc(h); err != nil {
		return nil, err
	}

	return app, nil
}

// AuthsApp type
type AuthsApp struct {
	handler           *handler.Handler
	AuthenticationSvc *service.AuthenticationSvc
}
