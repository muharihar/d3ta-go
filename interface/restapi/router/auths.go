package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
)

// SetAuths set Auths Router
func SetAuths(eg *echo.Group, f *feature.FAuths) {

	gc := eg.Group("/auths")

	gc.POST("/register", f.RegisterUser)
	gc.GET("/registration/activate/:activationCode/:format", f.ActivateRegistration)
	gc.POST("/login", f.Login)
	gc.POST("/login-app", f.LoginApp)
}
