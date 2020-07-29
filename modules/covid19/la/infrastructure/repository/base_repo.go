package repository

import (
	"fmt"
	"strings"

	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infC19Adp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19"
	"github.com/muharihar/d3ta-go/system/handler"
)

// BaseRepo represent BaseRepo
type BaseRepo struct {
	handler     *handler.Handler
	c19Adapters map[string]C19RepoAdapter
}

// C19RepoAdapter type
type C19RepoAdapter struct {
	Provider string
	Adapter  infC19Adp.ICovid19Adapter
}

// SetHandler set Handler
func (br *BaseRepo) SetHandler(h *handler.Handler) {
	br.handler = h
}

// GetHandler set Handler
func (br *BaseRepo) GetHandler() *handler.Handler {
	return br.handler
}

// AddAdapter is a function to Add Adapter Repository
func (br *BaseRepo) AddAdapter(alias, provider string, adapter infC19Adp.ICovid19Adapter) {
	if br.c19Adapters == nil {
		br.c19Adapters = make(map[string]C19RepoAdapter)
	}
	tmp := C19RepoAdapter{Provider: provider, Adapter: adapter}
	br.c19Adapters[alias] = tmp
}

// GetAdapters is a function to get availables adapters
func (br *BaseRepo) GetAdapters() map[string]C19RepoAdapter {
	return br.c19Adapters
}

// SelectedAdaptersByProviders is a funtion to select selected adapters from selected providers
func (br *BaseRepo) SelectedAdaptersByProviders(providers domSchema.ProviderList) map[string]C19RepoAdapter {
	tmp := make(map[string]C19RepoAdapter)

	for _, prov := range providers {
		provCode := strings.ToLower(prov.Code)
		// fmt.Println(k, " ---- ", provCode)
		if provCode == "_all_" {
			tmp := make(map[string]C19RepoAdapter)

			for k, p := range br.c19Adapters {
				if k != "_default_" {
					tmp[p.Provider] = p
				}
			}

			return tmp
		}
		if vAdp, ok := br.c19Adapters[provCode]; ok {
			tmp[vAdp.Provider] = vAdp
		}
	}

	if len(tmp) < 1 {
		tmp["_default_"] = br.c19Adapters["_default_"]
	}
	return tmp
}

// SelectAdapterByProvider select Adapter by Provider Code
func (br *BaseRepo) SelectAdapterByProvider(providerCode string) (*C19RepoAdapter, error) {
	a, ok := br.c19Adapters[strings.ToLower(providerCode)]
	if !ok {
		return nil, fmt.Errorf("Adapter Not found [%s]", providerCode)
	}
	return &a, nil
}
