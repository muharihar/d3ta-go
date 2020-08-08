package handler

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
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
	config          *config.Config
	dbGorms         map[string]*gorm.DB
	casbinEnforcers map[string]*casbin.Enforcer
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

// SetCasbinEnforcer set CasbinEnforcer
func (h *Handler) SetCasbinEnforcer(ceName string, ce *casbin.Enforcer) {
	if h.casbinEnforcers == nil {
		h.casbinEnforcers = make(map[string]*casbin.Enforcer)
	}
	h.casbinEnforcers[ceName] = ce
}

// GetCasbinEnforcer get CasbinEnforcer
func (h *Handler) GetCasbinEnforcer(ceName string) (*casbin.Enforcer, error) {
	ce, exist := h.casbinEnforcers[ceName]
	if !exist {
		err := fmt.Errorf("Casbin Enforcer Name '%s' Not Found", ceName)
		return nil, err
	}
	return ce, nil
}
