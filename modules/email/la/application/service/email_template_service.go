package service

import (
	"fmt"

	appDTO "github.com/muharihar/d3ta-go/modules/email/la/application/dto"
	domRepo "github.com/muharihar/d3ta-go/modules/email/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
	infRepo "github.com/muharihar/d3ta-go/modules/email/la/infrastructure/repository"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewEmailTemplateService new EmailTemplateService
func NewEmailTemplateService(h *handler.Handler) (*EmailTemplateService, error) {
	var err error

	svc := new(EmailTemplateService)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewEmailTemplateRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// EmailTemplateService type
type EmailTemplateService struct {
	BaseService
	repo domRepo.IEmailTemplateRepo
}

// ListAll list all email template
func (s *EmailTemplateService) ListAll(i identity.Identity) (*appDTO.ETListAllResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	res, err := s.repo.ListAll(i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETListAllResDTO)
	resDTO.Count = res.Count
	resDTO.Data = res.Data

	return resDTO, nil
}

// FindByCode find Email Template by Code
func (s *EmailTemplateService) FindByCode(req *appDTO.ETFindByCodeReqDTO, i identity.Identity) (*appDTO.ETFindByCodeResDTO, error) {
	// authorization
	if (i.CanAccessCurrentRequest() == false) && (i.CanAccess("", "system.module.email.template.findbycode", "READ", nil) == false) {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.ETFindByCodeRequest{
		Code: req.Code,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.FindByCode(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETFindByCodeResDTO)
	resDTO.Query = reqDom
	resDTO.Data = res.Data

	return resDTO, nil
}

// Create create new Email Template
func (s *EmailTemplateService) Create(req *appDTO.ETCreateReqDTO, i identity.Identity) (*appDTO.ETCreateResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.ETCreateRequest{
		Code:        req.Code,
		Name:        req.Name,
		IsActive:    req.IsActive,
		EmailFormat: req.EmailFormat,
		Template:    req.Template,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Create(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETCreateResDTO)
	resDTO.Code = res.Code
	resDTO.Version = res.Version

	return resDTO, nil
}

// Update update existing Email Template
func (s *EmailTemplateService) Update(req *appDTO.ETUpdateReqDTO, i identity.Identity) (*appDTO.ETUpdateResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.ETUpdateRequest{
		Keys: &domSchema.ETUpdateKeys{
			Code: req.Keys.Code,
		},
		Data: &domSchema.ETUpdateData{
			Name:        req.Data.Name,
			IsActive:    req.Data.IsActive,
			EmailFormat: req.Data.EmailFormat,
			Template:    req.Data.Template,
		},
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Update(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETUpdateResDTO)
	resDTO.Code = res.Code
	resDTO.Version = res.Version

	return resDTO, nil
}

// SetActive set existing Email Template active status
func (s *EmailTemplateService) SetActive(req *appDTO.ETSetActiveReqDTO, i identity.Identity) (*appDTO.ETSetActiveResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.ETSetActiveRequest{
		Keys: &domSchema.ETSetActiveKeys{
			Code: req.Keys.Code,
		},
		Data: &domSchema.ETSetActiveData{
			IsActive: req.Data.IsActive,
		},
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.SetActive(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETSetActiveResDTO)
	resDTO.Code = res.Code
	resDTO.IsActive = res.IsActive

	return resDTO, nil
}

// Delete delete existing Email Template with template versions
func (s *EmailTemplateService) Delete(req *appDTO.ETDeleteReqDTO, i identity.Identity) (*appDTO.ETDeleteResDTO, error) {
	// authorization
	if i.CanAccessCurrentRequest() == false {
		errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
			i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
		return nil, sysError.CustomForbiddenAccess(errMsg)
	}

	// request domain
	reqDom := domSchema.ETDeleteRequest{
		Code: req.Code,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Delete(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ETDeleteResDTO)
	resDTO.Query = res.Query
	resDTO.Data = res.Data

	return resDTO, nil
}
