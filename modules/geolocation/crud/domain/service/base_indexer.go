package service

import (
	"github.com/muharihar/d3ta-go/system/indexer"
)

// BaseIndexerSvc represent BaseIndexerSvc
type BaseIndexerSvc struct {
	BaseSvc
	index   string
	indexer *indexer.Indexer
}

// SetIndexer set Indexer
func (s *BaseIndexerSvc) SetIndexer(context, container, component string) (*indexer.Indexer, error) {
	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	ce, err := s.handler.GetIndexer(cfg.Indexers.DataIndexer.ConnectionName)
	if err != nil {
		return nil, err
	}
	ce.Context = context
	ce.Container = container
	ce.Component = component

	s.indexer = ce

	return ce, nil
}
