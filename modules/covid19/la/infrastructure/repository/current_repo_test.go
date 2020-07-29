package repository

import (
	"encoding/json"
	"testing"

	domRepo "github.com/muharihar/d3ta-go/modules/covid19/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

func newConfig(t *testing.T) (*config.Config, error) {

	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newRepoIdent(t *testing.T) (domRepo.ICurrentRepo, identity.Identity, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, identity.Identity{}, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, identity.Identity{}, err
	}

	h.SetConfig(c)

	r, err := NewCurrentRepo(h)
	if err != nil {
		return nil, identity.Identity{}, err
	}

	i := newIdentity(h, t)

	return r, i, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}

	return i
}

func TestRepoDisplayCurrentDataByCountry(t *testing.T) {
	repo, i, err := newRepoIdent(t)
	if err != nil {
		t.Errorf("REPO: [%#v]", err.Error())
	}

	var req domSchema.DisplayCurrentDataByCountryRequest

	err = json.Unmarshal([]byte(`{ "countryCode": "ID", "providers": [  {"code": "_ALL_" } ] }`), &req)
	// err = json.Unmarshal([]byte(`{ "countryCode": "ID", "providers": [  {"code": "WHO" }, {"code": "WHO" } ] }`), &req)
	if err != nil {
		t.Errorf("Request: [%s]", err.Error())
	}

	res, err := repo.DisplayCurrentDataByCountry(&req, i)
	if err != nil {
		t.Errorf("Request: [%s]", err.Error())
	}

	if res == nil {
		t.Fail()
	}

	if res != nil {
		resJSON, err := json.Marshal(res)
		if err != nil {
			t.Errorf("json.Marshal: [%s]", err.Error())
		}
		t.Logf("Resp: %s", string(resJSON))
	}
}
