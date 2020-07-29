package service

import (
	"encoding/json"
	"testing"

	appDTO "github.com/muharihar/d3ta-go/modules/geolocation/crud/application/dto"
	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newCountrySvc(t *testing.T) (*CountrySvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabase(h); err != nil {
		return nil, nil, err
	}

	r, err := NewCountrySvc(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	if err := i.SetCasbinEnforcer("../../../../../conf/casbin/casbin_rbac_rest_model.conf"); err != nil {
		t.Errorf("SetCasbinEnforcer: %s", err.Error())
	}

	i.Claims.Username = "test.d3tago"
	i.Claims.AuthorityID = "group:admin"

	return i
}

func TestCountrySvc_ListAll(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/list-all"
	i.RequestInfo.RequestAction = "GET"

	resp, err := cSvc.ListAll(i)
	if err != nil {
		t.Errorf("ListAll: %s", err.Error())
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestCountrySvc_GetDetail(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	req := new(appDTO.GetCountryReqDTO)
	req.Code = "ID"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "GET"

	resp, err := cSvc.GetDetail(req, i)
	if err != nil {
		t.Errorf("GetDetail: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestCountrySvc_Add(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	req := new(appDTO.AddCountryReqDTO)
	req.Code = "XX"
	req.Name = "XX COUNTRY"
	req.ISO2Code = "XX"
	req.ISO3Code = ""
	req.WHORegion = "WPRO"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country"
	i.RequestInfo.RequestAction = "POST"

	resp, err := cSvc.Add(req, i)
	if err != nil {
		t.Errorf("Add: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestCountrySvc_Update(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	req := new(appDTO.UpdateCountryReqDTO)
	req.Keys = &appDTO.UpdateCountryKeysDTO{Code: "XX"}
	req.Data = &appDTO.UpdateCountryDataDTO{
		Name:      "XX COUNTRY UPDATED",
		ISO2Code:  "XX",
		ISO3Code:  "",
		WHORegion: "WPRO",
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "PUT"

	resp, err := cSvc.Update(req, i)
	if err != nil {
		t.Errorf("Update: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestCountrySvc_Delete(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	req := new(appDTO.DeleteCountryReqDTO)
	req.Code = "XX"

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "DELETE"

	resp, err := cSvc.Delete(req, i)
	if err != nil {
		t.Errorf("Delete: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}
