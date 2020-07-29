package handler

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"

	"github.com/muharihar/d3ta-go/system/config"
)

// NewHandler new Handler
func NewHandler() (*Handler, error) {
	h := new(Handler)

	h.dbGorms = make(map[string]*gorm.DB)

	return h, nil
}

// Handler represent Handler
type Handler struct {
	config  *config.Config
	dbGorms map[string]*gorm.DB
}

// SetConfig set Config
func (h *Handler) SetConfig(config *config.Config) {
	h.config = config
}

// GetConfig get Config
func (h *Handler) GetConfig() (*config.Config, error) {
	if h.config == nil {
		return nil, fmt.Errorf("ERROR: [%s]", "Configuration Does Not Exist")
	}
	return h.config, nil
}

// SetGormDB set GORM database connection by connection name
func (h *Handler) SetGormDB(conName string, dbCon *gorm.DB) {
	if h.dbGorms == nil {
		h.dbGorms = make(map[string]*gorm.DB)
	}
	h.dbGorms[conName] = dbCon
}

// GetGormDB get GORM database connection by connection name
func (h *Handler) GetGormDB(conName string) (*gorm.DB, error) {
	db, exist := h.dbGorms[conName]
	if !exist {
		err := fmt.Errorf("DB Connection Name '%s' Not Found", conName)
		return nil, err
	}

	return db, nil
}
