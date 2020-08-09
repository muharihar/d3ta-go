package feature

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"

	response "github.com/muharihar/d3ta-go/interface/restapi/response"
	sysErr "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// BaseFeature represent BaseFeature
type BaseFeature struct {
	handler *handler.Handler
}

// GetHandler get superHandler
func (f *BaseFeature) GetHandler() *handler.Handler {
	return f.handler
}

func (f *BaseFeature) translateErrorMessage(err error, c echo.Context) error {
	// fmt.Printf("Type: %T, Value: %#v", err, err)
	switch fmt.Sprintf("%T", err) {
	case "*echo.HTTPError":
		tmp := err.(*echo.HTTPError)

		data := map[string]interface{}{
			"errorType": "validation", "validationType": "server.HTTPError", "validations": tmp.Message.(string),
		}
		return response.FailDetailedwithCode(tmp.Code, data, "validation failed", c)

	case "validation.Errors":
		if e, ok := err.(validation.InternalError); ok {
			return e
		}

		data := map[string]interface{}{
			"errorType": "validation", "validationType": "validation.Errors", "validations": string(err.Error()),
		}
		return response.FailDetailedwithCode(http.StatusBadRequest, data, "validation failed", c)

	case "*error.SystemError":
		tmp := err.(*sysErr.SystemError)
		data := map[string]interface{}{
			"errorType": "SystemError", "errorMessage": tmp.Error(),
		}
		return response.FailDetailedwithCode(tmp.StatusCode, data, "Operation failed", c)

	default:
		data := map[string]interface{}{
			"type":  fmt.Sprintf("%T", err),
			"error": string(err.Error()),
		}
		return response.FailDetailedwithCode(http.StatusInternalServerError, data, "Operation failed with default error message", c)
	}
}

// SetIdentity set Identity
func (f *BaseFeature) SetIdentity(c echo.Context) (identity.Identity, error) {

	token := c.Get("identity.token.jwt")
	if token == nil {
		token = ""
	}

	claims := c.Get("identity.token.jwt.claims")
	if claims == nil {
		claims = &identity.JWTCustomClaims{}
	}

	i, err := identity.NewIdentity(identity.DefaultIdentity, identity.TokenJWT, token.(string), claims.(*identity.JWTCustomClaims), c, f.handler)

	return i, err
}

func (f *BaseFeature) inTestMode() bool {
	return strings.HasSuffix(os.Args[0], ".test")

}
