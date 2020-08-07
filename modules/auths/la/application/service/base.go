package service

import (
	"github.com/muharihar/d3ta-go/system/context"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// BaseService Type
type BaseService struct {
	handler        *handler.Handler
	systemIdentity identity.Identity
}

func (b *BaseService) initBaseService() error {
	// init system identity
	if err := b.initSystemIdentity(); err != nil {
		return err
	}
	return nil
}

func (b *BaseService) initSystemIdentity() error {
	j, err := identity.NewJWT(b.handler)
	if err != nil {
		return err
	}
	claims, token, _, err := j.GenerateSystemToken()
	if err != nil {
		return err
	}
	if b.systemIdentity, err = identity.NewIdentity(identity.TokenJWT, token, claims, context.NewCtx(context.SystemCtx), b.handler); err != nil {
		return err
	}
	return nil
}
