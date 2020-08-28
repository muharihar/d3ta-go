package service

import (
	"testing"
	"time"

	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func newCountryIndexerSvc(t *testing.T) (*CountryIndexerSvc, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, err
	}

	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		return nil, err
	}

	r, err := NewCountryIndexerSvc(h)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func TestCountryIndexerSvc_Methods(t *testing.T) {
	idx, err := newCountryIndexerSvc(t)
	if err != nil {
		t.Errorf("Error while creating CountryIndexerSvc: newCountryIndexerSvc() [%s]", err.Error())
		return
	}

	countries := []*domSchema.Country{
		{ID: -134, Code: "11", Name: "Test Country 11", ISO2Code: "Z1", ISO3Code: "ZZ1", WHORegion: "REGION"},
		{ID: -135, Code: "22", Name: "Test Country 22", ISO2Code: "Z2", ISO3Code: "ZZ2", WHORegion: "REGION"},
	}

	t.Run("CountryIndexerSvc.Refresh", func(t *testing.T) {
		if err := idx.Refresh(countries); err != nil {
			t.Errorf("Error while refreshing country indexer: Refresh() [%s]", err.Error())
		}
	})
	// give time for elastic search
	time.Sleep(3 * time.Second)

	t.Run("CountryIndexerSvc.Search", func(t *testing.T) {
		res, err := idx.Search(&domSchema.SearchCountryIndexerRequest{Name: "TEST"})
		if err != nil {
			t.Errorf("Error while searcing country indexer: Search() [%s]", err.Error())
		}
		// t.Log("Response: ", string(res.ToJSON()))
		if res.Data == nil {
			t.Errorf("Should be not null/nil")
		}
	})

	t.Run("CountryIndexerSvc.Create", func(t *testing.T) {
		if err := idx.Create(countries[0]); err == nil {
			t.Errorf("Shoud be error: Create() [Index Document already exist*]")
		}
	})

	t.Run("CountryIndexerSvc.Update", func(t *testing.T) {
		if err := idx.Update(countries[0]); err != nil {
			t.Errorf("Error while updating country doc indexer: Update() [%s]", err.Error())
		}
	})

	t.Run("CountryIndexerSvc.Delete", func(t *testing.T) {
		for k, country := range countries {
			t.Logf("Delete Data: %d", k)
			if err := idx.Delete(country); err != nil {
				t.Errorf("Error while deleting country doc indexer: Delete() [%s]", err.Error())
			}
		}
	})
}
