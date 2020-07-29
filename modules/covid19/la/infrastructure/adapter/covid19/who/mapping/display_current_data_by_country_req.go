package mapping

import (
	con19type "github.com/muharihar/d3ta-go/connector/covid19/who/types"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
)

// MapDisplayCurrentDataByCountryReq mapping DisplayCurrentDataByCountryReq
func MapDisplayCurrentDataByCountryReq(req *domSchema.DisplayCurrentDataByCountryRequest) (*con19type.GetCountryRequest, error) {

	reqCon := new(con19type.GetCountryRequest)
	reqCon.CountryCode = req.CountryCode

	return reqCon, nil
}
