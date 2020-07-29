package service

import (
	appDTO "github.com/muharihar/d3ta-go/modules/auths/la/application/dto"
	domRepo "github.com/muharihar/d3ta-go/modules/auths/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/auths/la/domain/schema"
	infRepo "github.com/muharihar/d3ta-go/modules/auths/la/infrastructure/repository"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewAuthenticationSvc new AuthenticationSvc
func NewAuthenticationSvc(h *handler.Handler) (*AuthenticationSvc, error) {
	var err error

	svc := new(AuthenticationSvc)
	svc.handler = h

	if svc.repo, err = infRepo.NewAuthenticationRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// AuthenticationSvc type
type AuthenticationSvc struct {
	BaseService
	repo domRepo.IAuthenticationRepo
}

// Register user
func (s *AuthenticationSvc) Register(req *appDTO.RegisterReqDTO, i identity.Identity) (*appDTO.RegisterResDTO, error) {
	reqDom := domSchema.RegisterRequest{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		NickName:  req.NickName,
		Captcha:   req.Captcha,
		CaptchaID: req.CaptchaID,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Register(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.RegisterResDTO)
	resDTO.Email = res.Email

	return resDTO, nil
}

// ActivateRegistration activate user registration
func (s *AuthenticationSvc) ActivateRegistration(req *appDTO.ActivateRegistrationReqDTO, i identity.Identity) (*appDTO.ActivateRegistrationResDTO, error) {
	reqDom := domSchema.ActivateRegistrationRequest{
		ActivationCode: req.ActivationCode,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.ActivateRegistration(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.ActivateRegistrationResDTO)
	resDTO.Email = res.Email
	resDTO.DefaultRole = res.DefaultRole

	return resDTO, nil
}

// Login user
func (s *AuthenticationSvc) Login(req *appDTO.LoginReqDTO, i identity.Identity) (*appDTO.LoginResDTO, error) {
	reqDom := domSchema.LoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		Captcha:   req.Captcha,
		CaptchaID: req.CaptchaID,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Login(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.LoginResDTO)
	resDTO.TokenType = res.TokenType
	resDTO.Token = res.Token
	resDTO.ExpiredAt = res.ExpiredAt

	return resDTO, nil
}

// LoginApp login app
func (s *AuthenticationSvc) LoginApp(req *appDTO.LoginAppReqDTO, i identity.Identity) (*appDTO.LoginAppResDTO, error) {
	reqDom := domSchema.LoginAppRequest{
		ClientKey: req.ClientKey,
		SecretKey: req.SecretKey,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.LoginApp(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.LoginAppResDTO)
	resDTO.TokenType = res.TokenType
	resDTO.ClientAppCode = res.ClientAppCode
	resDTO.ClientAppName = res.ClientAppName
	resDTO.Token = res.Token
	resDTO.ExpiredAt = res.ExpiredAt

	return resDTO, nil
}
