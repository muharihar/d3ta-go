package repository

import (
	"fmt"
	"sync"

	domRepo "github.com/muharihar/d3ta-go/modules/covid19/la/domain/repository"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
	infC19goidAdp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19/covid19goid"
	infC19WHOAdp "github.com/muharihar/d3ta-go/modules/covid19/la/infrastructure/adapter/covid19/who"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/identity"
)

// NewCurrentRepo new CurrentRepo
func NewCurrentRepo(h *handler.Handler) (domRepo.ICurrentRepo, error) {

	repo := new(CurrentRepo)
	repo.handler = h

	adpWHO, adpWHOInfo, err := infC19WHOAdp.NewCovid19WHOAdapter(h)
	if err != nil {
		return nil, err
	}
	repo.AddAdapter(adpWHOInfo.Code, adpWHOInfo.Code, adpWHO)

	adpC19, adpC19Info, err := infC19goidAdp.NewCovid19goidAdapter(h)
	if err != nil {
		return nil, err
	}
	repo.AddAdapter(adpC19Info.Code, adpC19Info.Code, adpC19)

	repo.AddAdapter("_default_", adpWHOInfo.Code, adpWHO)
	return repo, nil
}

// CurrentRepo implement domRepo.CurrentRepo
type CurrentRepo struct {
	BaseRepo
}

// DisplayCurrentDataByCountry display CurrentDataByCountry
func (r *CurrentRepo) DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest, i identity.Identity) (*domSchema.DisplayCurrentDataByCountryResponse, error) {

	resp := new(domSchema.DisplayCurrentDataByCountryResponse)
	resp.Query = req

	var tmpData domSchema.TotalCountryProviderDataList
	selectedAdapters := r.SelectedAdaptersByProviders(req.Providers)

	//-->
	wg := sync.WaitGroup{}
	wg.Add(len(selectedAdapters))
	//<--

	for _, prv := range selectedAdapters {
		// fmt.Println(prv.Provider)
		fn := func(prv C19RepoAdapter) {
			//-->
			defer wg.Done()
			//<--
			tmp := new(domSchema.TotalCountryProviderData)
			tmp.Provider = prv.Provider

			adp := prv.Adapter
			if adp != nil {
				resAdp, err := adp.DisplayCurrentDataByCountry(req)
				if err != nil {
					tmp.Information = fmt.Sprintf("ERROR: %s", err.Error())
				} else {
					tmp.Information = resAdp.Information
					tmp.Data = resAdp.Data
				}
			}

			tmpData = append(tmpData, tmp)
		}
		// fn(prv)
		//-->
		go fn(prv)
		//<--
	}
	//-->
	wg.Wait()
	//<--

	resp.Data = &tmpData

	return resp, nil
}
