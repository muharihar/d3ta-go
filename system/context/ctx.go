package context

import (
	"fmt"

	"github.com/muharihar/d3ta-go/system/utils"
)

// Ctx is subject to be refactored

// CtxType type
type CtxType string

const (
	// SystemCtx const
	SystemCtx CtxType = "system.context"
)

// NewCtx new Ctx
func NewCtx(t CtxType) *Ctx {
	c := Ctx{
		ctxType: t,
	}
	return &c
}

// Ctx type
type Ctx struct {
	ctxType CtxType
}

// CtxRequest type
type CtxRequest struct {
	userAgent  string
	hostName   string
	remoteAddr string
	requestURI string
	method     string
}

// UserAgent get UserAgent
func (cr *CtxRequest) UserAgent() string {
	return cr.userAgent
}

// HostName get hostname
func (cr *CtxRequest) HostName() string {
	if cr.hostName == "" {
		cr.hostName = utils.GetHostName()
	}
	return cr.hostName
}

// RemoteAddr get hostname
func (cr *CtxRequest) RemoteAddr() string {
	return cr.remoteAddr
}

// RequestURI get requestURI
func (cr *CtxRequest) RequestURI() string {
	return cr.requestURI
}

// Method get method
func (cr *CtxRequest) Method() string {
	return cr.method
}

// Request get current request
func (c *Ctx) Request() *CtxRequest {
	cr := CtxRequest{}
	switch c.ctxType {
	case SystemCtx:
		cr.userAgent = "system.d3tago"
		cr.remoteAddr = fmt.Sprintf("system@%s", utils.GetCurrentIP())
		cr.requestURI = fmt.Sprintf("system@%s", utils.GetHostName())
		cr.method = "system.method"
	default:
		cr.userAgent = "default.d3tago"
		cr.remoteAddr = fmt.Sprintf("default@%s", utils.GetCurrentIP())
		cr.requestURI = fmt.Sprintf("default@%s", utils.GetHostName())
		cr.method = "default.method"
	}
	return &cr
}

// RealIP get real IP Address
func (c *Ctx) RealIP() string {
	return utils.GetCurrentIP()
}
