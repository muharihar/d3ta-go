package entity

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		t.Errorf("LoadAllDatabase: %s", err.Error())
	}

	return h, nil
}

func TestMigratoin(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
		return
	}

	cfg, err := h.GetConfig()
	if err != nil {
		t.Errorf("GetConfig: %s", err.Error())
		return
	}

	db, err := h.GetGormDB(cfg.Databases.EmailDB.ConnectionName)
	if err != nil {
		t.Errorf("GetGormDB: %s", err.Error())
		return
	}

	if err := db.AutoMigrate(&Email{}, &EmailTemplate{}, &EmailTemplateVersion{}); err != nil {
		t.Errorf("AutoMigrate: %s", err.Error())
	}
}
