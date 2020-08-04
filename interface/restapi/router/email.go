package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/feature"
	internalMiddleware "github.com/muharihar/d3ta-go/interface/restapi/middleware"
)

// SetEmail set Email Router
func SetEmail(eg *echo.Group, f *feature.FEmail) {

	gc := eg.Group("/email")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("/send", f.SendEmail)

	gc.GET("/templates/list-all", f.ListAllEmailTemplate)
	gc.GET("/template/:code", f.FindEmailTemplateByCode)
	gc.POST("/template", f.CreateEmailTemplate)
	gc.PUT("/template/update/:code", f.UpdateEmailTemplate)
	gc.PUT("/template/set-active/:code", f.SetActiveEmailTemplate)
	gc.DELETE("/template/:code", f.DeleteEmailTemplate)
}
