package initialize

import (
	"testing"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../conf")
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

	return h, nil
}

func TestLoadAllDatabase(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
	}

	if err = LoadAllDatabase(h); err != nil {
		t.Errorf("newHandler: %s", err.Error())
	}
}
