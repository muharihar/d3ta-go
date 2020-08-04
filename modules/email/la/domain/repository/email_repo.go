package repository

import (
	"github.com/muharihar/d3ta-go/modules/email/la/domain/schema"
	"github.com/muharihar/d3ta-go/system/identity"
)

// IEmailRepo interface
type IEmailRepo interface {
	Send(req *schema.SendEmailRequest, i identity.Identity) (*schema.SendEmailResponse, error)
}
