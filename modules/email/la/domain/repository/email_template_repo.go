package repository

import (
	schema "github.com/muharihar/d3ta-go/modules/email/la/domain/schema/email_template"
	"github.com/muharihar/d3ta-go/system/identity"
)

// IEmailTemplateRepo interface
type IEmailTemplateRepo interface {
	ListAll(i identity.Identity) (*schema.ETListAllResponse, error)
	FindByCode(req *schema.ETFindByCodeRequest, i identity.Identity) (*schema.ETFindByCodeResponse, error)
	Create(req *schema.ETCreateRequest, i identity.Identity) (*schema.ETCreateResponse, error)
	Update(req *schema.ETUpdateRequest, i identity.Identity) (*schema.ETUpdateResponse, error)
	SetActive(req *schema.ETSetActiveRequest, i identity.Identity) (*schema.ETSetActiveResponse, error)
	Delete(req *schema.ETDeleteRequest, i identity.Identity) (*schema.ETDeleteResponse, error)
}
