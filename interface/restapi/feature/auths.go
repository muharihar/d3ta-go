package feature

import (
	"fmt"
	"net/http"
	"strings"

	captcha "github.com/muharihar/d3ta-go/interface/restapi/feature/captcha"

	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/response"
	"github.com/muharihar/d3ta-go/modules/auths/la/application"
	"github.com/muharihar/d3ta-go/modules/auths/la/application/dto"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewFAuths new  FAuths
func NewFAuths(h *handler.Handler) (*FAuths, error) {
	var err error

	f := new(FAuths)
	f.handler = h

	if f.appAuths, err = application.NewAuthsApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FAuths feature Auths
type FAuths struct {
	BaseFeature
	appAuths *application.AuthsApp
}

// RegisterUser register user
func (f *FAuths) RegisterUser(c echo.Context) error {

	req := new(dto.RegisterReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	if !f.inTestMode() {
		decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
		if err != nil {
			return response.FailWithMessage(err.Error(), c)
		}

		if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
			return response.FailWithMessage("Captcha verification code error", c)
		}
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.Register(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// ActivateRegistration activater user registration
func (f *FAuths) ActivateRegistration(c echo.Context) error {

	//params
	format := strings.ToLower(c.Param("format"))

	req := new(dto.ActivateRegistrationReqDTO)
	req.ActivationCode = c.Param("activationCode")

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.ActivateRegistration(req, i)
	if err != nil {
		if format == "html" {
			data := map[string]interface{}{
				"message": err.Error(),
			}
			return c.Render(http.StatusBadRequest, "auths/activate.registration", data)
		}

		return f.translateErrorMessage(err, c)
	}

	if format == "html" {
		data := map[string]interface{}{
			"message": fmt.Sprintf("Your user [%s] are now active", resp.Email),
		}
		return c.Render(http.StatusOK, "auths/activate.registration", data)
	}

	return response.OkWithData(resp, c)
}

// Login user Login
func (f *FAuths) Login(c echo.Context) error {

	req := new(dto.LoginReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	if !f.inTestMode() {
		decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
		if err != nil {
			return response.FailWithMessage(err.Error(), c)
		}

		if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
			return response.FailWithMessage("Captcha verification code error", c)
		}
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.Login(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// LoginApp login client app
func (f *FAuths) LoginApp(c echo.Context) error {

	req := new(dto.LoginAppReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.LoginApp(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}
