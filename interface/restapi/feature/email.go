package feature

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/response"
	appEmail "github.com/muharihar/d3ta-go/modules/email/la/application"
	appEmailDTO "github.com/muharihar/d3ta-go/modules/email/la/application/dto"
	"github.com/muharihar/d3ta-go/system/handler"
)

// NewFEmail new FEmail
func NewFEmail(h *handler.Handler) (*FEmail, error) {
	var err error

	f := new(FEmail)
	f.handler = h

	if f.appEmail, err = appEmail.NewEmailApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FEmail represent Email Feature
type FEmail struct {
	BaseFeature
	appEmail *appEmail.EmailApp
}

// ListAllEmailTemplate list all EmailTemplate
func (f *FEmail) ListAllEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	resp, err := f.appEmail.EmailTemplateSvc.ListAll(i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// FindEmailTemplateByCode find EmailTemplateByCode
func (f *FEmail) FindEmailTemplateByCode(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appEmailDTO.ETFindByCodeReqDTO)
	req.Code = code

	resp, err := f.appEmail.EmailTemplateSvc.FindByCode(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// CreateEmailTemplate create EmailTemplate
func (f *FEmail) CreateEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appEmailDTO.ETCreateReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appEmail.EmailTemplateSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// UpdateEmailTemplate update existing EmailTemplate
func (f *FEmail) UpdateEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	reqKeys := new(appEmailDTO.ETUpdateKeysDTO)
	reqKeys.Code = code

	reqData := new(appEmailDTO.ETUpdateDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.translateErrorMessage(err, c)
	}

	req := &appEmailDTO.ETUpdateReqDTO{
		Keys: reqKeys,
		Data: reqData,
	}

	resp, err := f.appEmail.EmailTemplateSvc.Update(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// SetActiveEmailTemplate set existing EmailTemplate active status
func (f *FEmail) SetActiveEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	reqKeys := new(appEmailDTO.ETSetActiveKeysDTO)
	reqKeys.Code = code

	reqData := new(appEmailDTO.ETSetActiveDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.translateErrorMessage(err, c)
	}

	req := &appEmailDTO.ETSetActiveReqDTO{
		Keys: reqKeys,
		Data: reqData,
	}

	resp, err := f.appEmail.EmailTemplateSvc.SetActive(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// DeleteEmailTemplate delete existing EmailTemplate with template version
func (f *FEmail) DeleteEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	req := new(appEmailDTO.ETDeleteReqDTO)
	req.Code = code

	resp, err := f.appEmail.EmailTemplateSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// SendEmail send Email
func (f *FEmail) SendEmail(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appEmailDTO.SendEmailReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appEmail.EmailSvc.Send(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}
