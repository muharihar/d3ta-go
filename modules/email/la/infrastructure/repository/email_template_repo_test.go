package repository

import (
	"testing"

	"github.com/muharihar/d3ta-go/modules/email/la/domain/repository"
	schema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newEmailTemplateRepo(t *testing.T) (repository.IEmailTemplateRepo, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, nil, err
	}

	r, err := NewEmailTemplateRepo(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestEMailTemplateRepo_ListAll(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	i := newIdentity(h, t)

	res, err := repo.ListAll(i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.ListAll: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.ListAll: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_FindByCode(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	req := &schema.ETFindByCodeRequest{
		Code: "test.code",
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.FindByCode(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.FindByCode: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.FindByCode: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Create(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	req := &schema.ETCreateRequest{
		Code:     "test.code.02",
		Name:     "Template Name 02",
		IsActive: true,
		Template: &schema.ETCreateVersion{
			SubjectTpl: "Subject Template",
			BodyTpl:    `{{define "T"}}Body Template{{end}}`,
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Create(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Create: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.Create: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Update(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	req := &schema.ETUpdateRequest{
		Keys: &schema.ETUpdateKeys{
			Code: "test.code.02",
		},
		Data: &schema.ETUpdateData{
			Name:     "Template Name 02 Updated",
			IsActive: true,
			Template: &schema.ETUpdateVersion{
				SubjectTpl: "Subject Template Updated",
				BodyTpl:    `{{define "T"}}Body Template Updated{{end}}`,
			},
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Update(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Update: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.Update: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_SetActive(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	req := &schema.ETSetActiveRequest{
		Keys: &schema.ETSetActiveKeys{
			Code: "test.code.02",
		},
		Data: &schema.ETSetActiveData{
			IsActive: true,
		},
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.SetActive(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.SetActive: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.SetActive: %s", string(respJSON))
	}
}

func TestEMailTemplateRepo_Delete(t *testing.T) {
	repo, h, err := newEmailTemplateRepo(t)
	if err != nil {
		t.Errorf("Error.newEmailTemplateRepo: %s", err.Error())
	}

	req := &schema.ETDeleteRequest{
		Code: "test.code.02",
	}

	if err := req.Validate(); err != nil {
		t.Errorf("Error.Req.Validate: %s", err.Error())
		return
	}

	i := newIdentity(h, t)

	res, err := repo.Delete(req, i)
	if err != nil {
		t.Errorf("Error.EmailTemplateRepo.Delete: %s", err.Error())
		return
	}

	if res != nil {
		respJSON := res.ToJSON()
		t.Logf("Resp.EmailTemplateRepo.Delete: %s", string(respJSON))
	}
}
