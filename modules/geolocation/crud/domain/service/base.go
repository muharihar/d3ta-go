package service

import "github.com/muharihar/d3ta-go/system/handler"

// BaseSvc represent BaseSvc
type BaseSvc struct {
	handler          *handler.Handler
	dbConnectionName string
}

// SetHandler set Handler
func (s *BaseSvc) SetHandler(h *handler.Handler) {
	s.handler = h
}

// GetHandler set Handler
func (s *BaseSvc) GetHandler() *handler.Handler {
	return s.handler
}

// SetDBConnectionName set DBConnectionName
func (s *BaseSvc) SetDBConnectionName(v string) {
	s.dbConnectionName = v
}
