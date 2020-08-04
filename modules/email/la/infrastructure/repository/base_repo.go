package repository

import "github.com/muharihar/d3ta-go/system/handler"

// BaseRepo type
type BaseRepo struct {
	handler          *handler.Handler
	dbConnectionName string
}

// SetHandler set Handler
func (r *BaseRepo) SetHandler(h *handler.Handler) {
	r.handler = h
}

// SetDBConnectionName set DBConnectionName
func (r *BaseRepo) SetDBConnectionName(v string) {
	r.dbConnectionName = v
}
