package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	domSchema "github.com/muharihar/d3ta-go/modules/geolocation/crud/domain/schema/country"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/indexer/schema"
)

// NewCountryIndexerSvc new CountryIndexerSvc
func NewCountryIndexerSvc(h *handler.Handler) (*CountryIndexerSvc, error) {
	svc := new(CountryIndexerSvc)
	svc.SetHandler(h)

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	svc.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	if svc.indexer, err = svc.SetIndexer("module", "mdm", "geolocation"); err != nil {
		return nil, err
	}
	svc.index = "countries"

	if err = svc.init(); err != nil {
		return nil, err
	}

	return svc, nil
}

// CountryIndexerSvc represent CountryIndexerSvc
type CountryIndexerSvc struct {
	BaseIndexerSvc
}

type esSearchResponse struct {
	schema.ESIndexSearchResponse
	Hits *esHits `json:"hits"`
}

type esHits struct {
	schema.ESHits
	Hits []*esHit `json:"hits"`
}

type esHit struct {
	schema.ESHit
	Source *domSchema.Country `json:"_source"`
}

func (s *CountryIndexerSvc) init() error {
	exist, err := s.indexer.IndexExist([]string{s.index})
	if err != nil {
		return err
	}
	if !exist {
		_mapping := `
{
    "mappings": {
        "properties": {
            "ID": {
                "type": "long"
			},
			"code": {
                "type": "text",
                "fields": { 
					"keyword": { "type":"keyword", "ignore_above":256 }
                }
            },
            "name": {
                "type": "text",
                "fields": {
                    "keyword": { "type":"keyword", "ignore_above":256 }
                }
            },
            "ISO2Code": {
                "type": "text",
                "fields": {
                    "keyword": { "type":"keyword", "ignore_above":256 }
                }
            },
            "ISO3Code": {
                "type": "text",
                "fields": {
                    "keyword": { "type":"keyword", "ignore_above":256 }
                }
            },
            "WHORegion": {
                "type": "text",
                "fields": {
                    "keyword": { "type":"keyword", "ignore_above":256 }
                }
            }
        }
    }
}`
		if err := s.indexer.CreateIndex(s.index, strings.NewReader(_mapping)); err != nil {
			return err
		}
	}
	return nil
}

// Refresh country index
func (s *CountryIndexerSvc) Refresh(countries []*domSchema.Country) error {
	for _, country := range countries {
		exist, err := s.indexer.DocExist(s.index, strconv.FormatInt(country.ID, 10))
		if err != nil {
			continue
		}
		if exist {
			if err := s.Update(country); err != nil {
				return err
			}
		} else {
			if err := s.Create(country); err != nil {
				return err
			}
		}
	}
	return nil
}

// Search country index
func (s *CountryIndexerSvc) Search(req *domSchema.SearchCountryIndexerRequest) (*domSchema.SearchCountryIndexerResponse, error) {
	// query
	_query := fmt.Sprintf(`
{
    "size": %d,
    "query": {
        "wildcard": { "name": "*%s*" }
    },
    "sort": [
        { "name.keyword": { "order": "asc", "unmapped_type": "boolean" } }
    ]
}`, 25, req.Name)

	ri, err := s.indexer.SearchIndexDoc(s.index, strings.NewReader(_query), 25, false)
	if err != nil {
		return nil, err
	}

	var resES esSearchResponse
	if err := json.Unmarshal(ri, &resES); err != nil {
		return nil, err
	}

	// response
	var listCountry []*domSchema.Country
	for _, rec := range resES.Hits.Hits {
		tmp := new(domSchema.Country)

		tmp.ID = rec.Source.ID
		tmp.Code = rec.Source.Code
		tmp.Name = rec.Source.Name
		tmp.ISO2Code = rec.Source.ISO2Code
		tmp.ISO3Code = rec.Source.ISO3Code
		tmp.WHORegion = rec.Source.WHORegion

		listCountry = append(listCountry, tmp)
	}

	res := new(domSchema.SearchCountryIndexerResponse)
	res.Query = req
	res.Data = listCountry

	return res, nil
}

// Create new country document
func (s *CountryIndexerSvc) Create(country *domSchema.Country) error {
	if err := s.indexer.CreateDoc(s.index, strconv.FormatInt(country.ID, 10), strings.NewReader(string(country.ToJSON()))); err != nil {
		return err
	}
	return nil
}

// Update existing country document
func (s *CountryIndexerSvc) Update(country *domSchema.Country) error {
	update := fmt.Sprintf(`{ "doc": %s }`, string(country.ToJSON()))
	if err := s.indexer.UpdateDoc(s.index, strconv.FormatInt(country.ID, 10), strings.NewReader(update)); err != nil {
		return err
	}
	return nil
}

// Delete existing country document
func (s *CountryIndexerSvc) Delete(country *domSchema.Country) error {
	if err := s.indexer.DeleteDoc(s.index, strconv.FormatInt(country.ID, 10)); err != nil {
		return err
	}
	return nil
}
