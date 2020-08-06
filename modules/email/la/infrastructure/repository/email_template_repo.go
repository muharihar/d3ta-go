package repository

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blang/semver/v4"
	domEntity "github.com/muharihar/d3ta-go/modules/email/la/domain/entity"
	domRepo "github.com/muharihar/d3ta-go/modules/email/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
	sysError "github.com/muharihar/d3ta-go/system/error"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
	"github.com/muharihar/d3ta-go/system/utils"
)

// NewEmailTemplateRepo new EmailTemplateRepo implement IEmailTemplateRepo
func NewEmailTemplateRepo(h *handler.Handler) (domRepo.IEmailTemplateRepo, error) {

	repo := new(EmailTemplateRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.EmailDB.ConnectionName)

	return repo, nil
}

// EmailTemplateRepo type implement IEmailTemplateRepo
type EmailTemplateRepo struct {
	BaseRepo
}

// ListAll list all email template
func (r *EmailTemplateRepo) ListAll(i identity.Identity) (*domSchema.ETListAllResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var emailTplEtts []domEntity.EmailTemplate
	var count int64

	if err := dbCon.Order("id ASC").Find(&emailTplEtts).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	var listET []*domSchema.EmailTemplate
	for _, rec := range emailTplEtts {
		tmp := new(domSchema.EmailTemplate)

		tmp.ID = rec.ID
		tmp.UUID = rec.UUID
		tmp.Code = rec.Code
		tmp.Name = rec.Name
		tmp.IsActive = rec.IsActive
		tmp.EmailFormat = rec.EmailFormat
		tmp.DefaultVersionID = rec.DefaultVersionID

		listET = append(listET, tmp)
	}

	resp := new(domSchema.ETListAllResponse)
	resp.Count = count
	resp.Data = listET

	return resp, nil
}

// FindByCode find EmailTemplate by code
func (r *EmailTemplateRepo) FindByCode(req *domSchema.ETFindByCodeRequest, i identity.Identity) (*domSchema.ETFindByCodeResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var emailTplEtt domEntity.EmailTemplate
	var emailTplVersionEtt domEntity.EmailTemplateVersion
	var count int64

	if err := dbCon.Where("et.code = ?", req.Code).Preload("EmailTemplate").Joins(fmt.Sprintf(`JOIN %s et ON et.default_version_id = %s.id`, emailTplEtt.TableName(), emailTplVersionEtt.TableName())).Find(&emailTplVersionEtt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Email Template Code")}
	}

	// check if template is active
	if emailTplVersionEtt.EmailTemplate.IsActive == false {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("In Active Email Template")}
	}

	// response
	resp := new(domSchema.ETFindByCodeResponse)
	et := domSchema.EmailTemplate{
		ID:               emailTplVersionEtt.EmailTemplate.ID,
		UUID:             emailTplVersionEtt.EmailTemplate.UUID,
		Code:             emailTplVersionEtt.EmailTemplate.Code,
		Name:             emailTplVersionEtt.EmailTemplate.Name,
		IsActive:         emailTplVersionEtt.EmailTemplate.IsActive,
		EmailFormat:      emailTplVersionEtt.EmailTemplate.EmailFormat,
		DefaultVersionID: emailTplVersionEtt.EmailTemplate.DefaultVersionID,
	}
	dtv := domSchema.EmailTemplateVersion{
		ID:              emailTplVersionEtt.ID,
		Version:         emailTplVersionEtt.Version,
		SubjectTpl:      emailTplVersionEtt.SubjectTpl,
		BodyTpl:         emailTplVersionEtt.BodyTpl,
		EmailTemplateID: emailTplVersionEtt.EmailTemplateID,
	}
	d := domSchema.ETFindByCodeData{
		EmailTemplate:          et,
		DefaultTemplateVersion: dtv,
	}
	resp.Data = d

	return resp, nil
}

// Create create new template
func (r *EmailTemplateRepo) Create(req *domSchema.ETCreateRequest, i identity.Identity) (*domSchema.ETCreateResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// start transaction
	// -->
	tx := dbCon.Begin()
	defer func() {
		if r := recover(); r != nil {
			// <--
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// create email template
	emailTplEtt := domEntity.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        req.Code,
		Name:        req.Name,
		EmailFormat: req.EmailFormat,
		IsActive:    req.IsActive,
	}
	emailTplEtt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := tx.Create(&emailTplEtt).Error; err != nil {
		// <--
		tx.Rollback()
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// create email template version
	emailTplVerEtt := domEntity.EmailTemplateVersion{
		Version:         utils.GenSemVersion(""),
		SubjectTpl:      req.Template.SubjectTpl,
		BodyTpl:         req.Template.BodyTpl,
		EmailTemplateID: emailTplEtt.ID,
	}
	emailTplVerEtt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := tx.Create(&emailTplVerEtt).Error; err != nil {
		// <--
		tx.Rollback()
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// set default email template version
	emailTplEtt.DefaultVersionID = emailTplVerEtt.ID
	emailTplEtt.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := tx.Save(&emailTplEtt).Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if err := tx.Commit().Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	resp := new(domSchema.ETCreateResponse)
	resp.Code = req.Code
	resp.Version = emailTplVerEtt.Version

	return resp, nil
}

// Update update email template
func (r *EmailTemplateRepo) Update(req *domSchema.ETUpdateRequest, i identity.Identity) (*domSchema.ETUpdateResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var emailTplEtt domEntity.EmailTemplate
	var emailTplVersionEtt domEntity.EmailTemplateVersion
	var count int64

	if err := dbCon.Where("et.code = ?", req.Keys.Code).Preload("EmailTemplate").Joins(fmt.Sprintf(`JOIN %s et ON et.default_version_id = %s.id`, emailTplEtt.TableName(), emailTplVersionEtt.TableName())).Find(&emailTplVersionEtt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Email Template Code")}
	}
	lastVersion := emailTplVersionEtt.Version

	var maxEmailTplVerEtt domEntity.EmailTemplateVersion
	if err := dbCon.Where("email_template_id = ?", emailTplVersionEtt.EmailTemplateID).Order("version DESC").First(&maxEmailTplVerEtt).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	// set last version
	v1, _ := semver.Parse(lastVersion)
	v2, _ := semver.Parse(maxEmailTplVerEtt.Version)
	if v1.Compare(v2) == -1 {
		lastVersion = maxEmailTplVerEtt.Version
	}

	// start transaction
	// -->
	tx := dbCon.Begin()
	defer func() {
		if r := recover(); r != nil {
			// <--
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// create new version
	newEmailTplVerEtt := domEntity.EmailTemplateVersion{
		Version:         utils.GenSemVersion(lastVersion),
		SubjectTpl:      req.Data.Template.SubjectTpl,
		BodyTpl:         req.Data.Template.BodyTpl,
		EmailTemplateID: emailTplVersionEtt.EmailTemplateID,
	}
	newEmailTplVerEtt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)
	if err := tx.Create(&newEmailTplVerEtt).Error; err != nil {
		// <--
		tx.Rollback()
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// update and set default email template version
	emailTplVersionEtt.EmailTemplate.Name = req.Data.Name
	emailTplVersionEtt.EmailTemplate.IsActive = req.Data.IsActive
	emailTplVersionEtt.EmailTemplate.DefaultVersionID = newEmailTplVerEtt.ID
	emailTplVersionEtt.EmailTemplate.EmailFormat = req.Data.EmailFormat
	emailTplVersionEtt.EmailTemplate.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := tx.Save(&emailTplVersionEtt.EmailTemplate).Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if err := tx.Commit().Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	resp := new(domSchema.ETUpdateResponse)
	resp.Code = req.Keys.Code
	resp.Version = newEmailTplVerEtt.Version

	return resp, nil
}

// SetActive set Email Template Active Status
func (r *EmailTemplateRepo) SetActive(req *domSchema.ETSetActiveRequest, i identity.Identity) (*domSchema.ETSetActiveResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var emailTplEtt domEntity.EmailTemplate
	var count int64

	if err := dbCon.Where("code = ? ", req.Keys.Code).Find(&emailTplEtt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Email Template Code")}
	}

	// update data
	emailTplEtt.IsActive = req.Data.IsActive
	emailTplEtt.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&emailTplEtt).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	resp := new(domSchema.ETSetActiveResponse)
	resp.Code = req.Keys.Code
	resp.IsActive = req.Data.IsActive

	return resp, nil
}

// Delete delete Email Template & Template Version by code
func (r *EmailTemplateRepo) Delete(req *domSchema.ETDeleteRequest, i identity.Identity) (*domSchema.ETDeleteResponse, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var emailTplEtt domEntity.EmailTemplate
	var count int64

	if err := dbCon.Where("code = ?", req.Code).Find(&emailTplEtt).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Invalid Email Template Code")}
	}

	var listEmailTplVerEtt []domEntity.EmailTemplateVersion
	var countV int64
	if err := dbCon.Where("email_template_id = ?", emailTplEtt.ID).Find(&listEmailTplVerEtt).Count(&countV).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// start transaction
	// -->
	tx := dbCon.Begin()
	defer func() {
		if r := recover(); r != nil {
			// <--
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	emailTplEtt.DeletedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)
	if err := tx.Save(&emailTplEtt).Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}
	if err := tx.Delete(&emailTplEtt).Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// we can replace for-loop with batch delete operation
	// this (for-loop) use case, is just for proof of concept (master detail delete transaction)
	for _, etv := range listEmailTplVerEtt {
		etv.DeletedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)
		if err := tx.Save(&etv).Error; err != nil {
			// <--
			tx.Rollback()
			return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
		}
		if err := tx.Delete(&etv).Error; err != nil {
			// <--
			tx.Rollback()
			return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
		}
	}

	if err := tx.Commit().Error; err != nil {
		// <--
		tx.Rollback()
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	resp := new(domSchema.ETDeleteResponse)
	resp.Query = req
	resp.Data = &domSchema.ETDeleteResponseData{
		EmailTemplate: domSchema.EmailTemplate{
			ID:               emailTplEtt.ID,
			UUID:             emailTplEtt.UUID,
			Code:             emailTplEtt.Code,
			Name:             emailTplEtt.Name,
			IsActive:         emailTplEtt.IsActive,
			EmailFormat:      emailTplEtt.EmailFormat,
			DefaultVersionID: emailTplEtt.DefaultVersionID,
		},
		VersionCount: countV,
	}

	return resp, nil
}
