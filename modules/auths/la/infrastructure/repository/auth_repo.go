package repository

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	domEntity "github.com/muharihar/d3ta-go/modules/auths/la/domain/entity"
	domRepo "github.com/muharihar/d3ta-go/modules/auths/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/auths/la/domain/schema"
	"github.com/muharihar/d3ta-go/modules/auths/la/infrastructure/template"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
	"github.com/muharihar/d3ta-go/system/service"
	"github.com/muharihar/d3ta-go/system/utils"
)

// NewAuthenticationRepo new AuthenticationRepo implement IAuthenticationRepo
func NewAuthenticationRepo(h *handler.Handler) (domRepo.IAuthenticationRepo, error) {
	repo := new(AuthenticationRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.IdentityDB.ConnectionName)

	repo.smtp, err = service.NewSMTPSender(h)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// AuthenticationRepo type
type AuthenticationRepo struct {
	BaseRepo
	smtp *service.SMTPSender
}

// Register user
func (r *AuthenticationRepo) Register(req *domSchema.RegisterRequest, i identity.Identity) (*domSchema.RegisterResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var tmpUserRegEntt domEntity.TmpUserRegistration
	var userEntt domEntity.SysUser
	var count int64

	// check if username or email is already registered
	if err := dbCon.Where("(username = ?) OR (email = ?)", req.Username, req.Email).Find(&tmpUserRegEntt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count > 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: fmt.Errorf("Username or Email already registered")}
	}

	// check if username or email is alreday activated
	if err := dbCon.Where("(username = ?) OR (email = ?)", req.Email, req.Email).Find(&userEntt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count > 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: fmt.Errorf("Username or Email already activated")}
	}

	// register user
	tmpUserRegEntt.UUID = utils.GenerateUUID()
	tmpUserRegEntt.Username = req.Username
	tmpUserRegEntt.Password = utils.MD5([]byte(req.Password))
	tmpUserRegEntt.NickName = req.NickName
	tmpUserRegEntt.Email = strings.ToLower(req.Email)
	tmpUserRegEntt.IsActivated = false
	tmpUserRegEntt.ActivationCode = utils.GenerateRegistrationActivationCode()

	tmpUserRegEntt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Create(&tmpUserRegEntt).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// return response
	resp := new(domSchema.RegisterResponse)
	resp.Email = req.Email

	// send activation email (backgourd, feature: event sourcing, OnRegistered Event send Email)
	// if needed, we can move this part to application (service) layer
	go r.sendActivationCodeEmail(req.Email, req.NickName, tmpUserRegEntt.ActivationCode, "html", i.RequestInfo.Host)
	// if err := r.sendActivationCodeEmail(req.Email, req.NickName, tmpUserRegEntt.ActivationCode, "html"); err != nil {
	//	  panic(err)
	// }

	return resp, nil
}

// ActivateRegistration avtivate user registration
func (r *AuthenticationRepo) ActivateRegistration(req *domSchema.ActivateRegistrationRequest, i identity.Identity) (*domSchema.ActivateRegistrationResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var tmpUserRegEntt domEntity.TmpUserRegistration
	var userEntt domEntity.SysUser
	var count int64

	// check if activation code is exist
	if err := dbCon.Where(" activation_code = ? AND is_activated = 0", req.ActivationCode).Find(&tmpUserRegEntt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Activation Code")}
	}

	// check if username or email already activated
	if err := dbCon.Where("(username = ?) OR (email = ?)", tmpUserRegEntt.Username, tmpUserRegEntt.Email).Find(&userEntt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count > 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusAlreadyReported, Err: fmt.Errorf("Username or Email already activated")}
	}

	// activated to user
	userEntt.UUID = tmpUserRegEntt.UUID
	userEntt.Username = tmpUserRegEntt.Username
	userEntt.Password = tmpUserRegEntt.Password
	userEntt.NickName = tmpUserRegEntt.NickName
	userEntt.Email = strings.ToLower(tmpUserRegEntt.Email)
	userEntt.IsActive = true

	userEntt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	cfg, err := r.handler.GetConfig()
	if err != nil {
		return nil, err
	}
	userEntt.AuthorityID = cfg.IAM.Registration.DefaultAuthorityID // "group:default"

	if err := dbCon.Create(&userEntt).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// update status
	tmpUserRegEntt.IsActivated = true
	now := time.Now()
	tmpUserRegEntt.ActivatedAt = &now

	tmpUserRegEntt.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&tmpUserRegEntt).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// return response
	resp := new(domSchema.ActivateRegistrationResponse)
	resp.Email = userEntt.Email
	resp.DefaultRole = userEntt.AuthorityID

	// send activation result
	// if needed, we can move this part to application (service) layer
	go r.sendActivationResultEmail(userEntt.Email, userEntt.NickName)
	// if err := r.sendActivationResultEmail(userEntt.Email, userEntt.NickName); err != nil {
	//	  panic(err)
	// }

	return resp, nil
}

// Login user
func (r *AuthenticationRepo) Login(req *domSchema.LoginRequest, i identity.Identity) (*domSchema.LoginResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var userEntt domEntity.SysUser
	var count int64
	req.Password = utils.MD5([]byte(req.Password))

	if err := dbCon.Where(" username = ? AND password = ? ", req.Username, req.Password).Find(&userEntt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid username or password")}
	}

	// check user is active
	if !userEntt.IsActive {
		return nil, &sysError.SystemError{StatusCode: http.StatusNonAuthoritativeInfo, Err: fmt.Errorf("Inactive User")}
	}

	// return response
	resp := new(domSchema.LoginResponse)
	resp.TokenType = "JWT"
	resp.Token, resp.ExpiredAt, err = r.generateJWTToken(&userEntt)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LoginApp login app
func (r *AuthenticationRepo) LoginApp(req *domSchema.LoginAppRequest, i identity.Identity) (*domSchema.LoginAppResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var userClientApp domEntity.SysUserClientApps
	var count int64

	if err := dbCon.Where("client_key =? AND secret_key = ?", req.ClientKey, req.SecretKey).Preload("User").Find(&userClientApp).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Client Key or Secret Key")}
	}

	// check is active
	if !userClientApp.IsActive {
		return nil, fmt.Errorf("Inactive User Client Apps [%s]", userClientApp.ClientAppCode)
	}

	// check user is active
	if !userClientApp.User.IsActive {
		return nil, fmt.Errorf("Inactive User with Client App Code [%s]", userClientApp.ClientAppCode)
	}

	// return response
	resp := new(domSchema.LoginAppResponse)
	resp.TokenType = "JWT"
	resp.ClientAppCode = userClientApp.ClientAppCode
	resp.ClientAppName = userClientApp.ClientAppName
	resp.Token, resp.ExpiredAt, err = r.generateJWTToken(&userClientApp.User)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *AuthenticationRepo) generateJWTToken(data *domEntity.SysUser) (token string, expiredAt int64, err error) {
	j, err := identity.NewJWT(r.handler)
	if err != nil {
		return "", 0, err
	}

	claims := j.CreateCustomClaims(data.ID, data.UUID, data.Username, data.NickName, data.AuthorityID)

	token, expiredAt, err = j.GenerateToken(claims)

	return token, expiredAt, err
}

// if needed, we can move this part to application (service) layer
func (r *AuthenticationRepo) sendActivationCodeEmail(toEmail, toName string, code, responseFormat string, host string) error {
	cfg, err := r.handler.GetConfig()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(cfg.IAM.Registration.ActivationURL, host)

	activationURL := fmt.Sprintf("%s/%s/%s", url, code, responseFormat)

	body := template.UserRegisterActivationTemplate(toEmail, toName, activationURL)

	return r.smtp.SendEmail(toEmail, fmt.Sprintf("[%s] Account Activation", cfg.Applications.Name), body)
}

// if needed, we can move this part to application (service) layer
func (r *AuthenticationRepo) sendActivationResultEmail(toEmail, toName string) error {
	cfg, err := r.handler.GetConfig()
	if err != nil {
		return err
	}

	body := template.UserRegisterActivatedTemplate(toEmail, toName)

	return r.smtp.SendEmail(toEmail, fmt.Sprintf("[%s] Account Activation Success", cfg.Applications.Name), body)
}
