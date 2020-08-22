package service

import (
	"github.com/muharihar/d3ta-go/system/cacher"
	"github.com/muharihar/d3ta-go/system/handler"
)

// BaseSvc represent BaseSvc
type BaseSvc struct {
	handler *handler.Handler
}

// SetCacher set Cacher
func (b *BaseSvc) SetCacher(context, container, component string) (*cacher.Cacher, error) {
	cfg, err := b.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	ce, err := b.handler.GetCacher(cfg.Caches.TmpDataCache.ConnectionName)
	if err != nil {
		return nil, err
	}
	ce.Context = context
	ce.Container = container
	ce.Component = component

	return ce, nil
}
