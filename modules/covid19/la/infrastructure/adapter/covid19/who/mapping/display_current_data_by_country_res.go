package mapping

import (
	con19type "github.com/muharihar/d3ta-go/connector/covid19/who/types"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
)

// MapDisplayCurrentDataByCountryRes Mapping DisplayCurrentDataByCountryRes
func MapDisplayCurrentDataByCountryRes(res *con19type.GetCountryResponse) (*domSchema.TotalCountryProviderData, error) {

	resDom := new(domSchema.TotalCountryProviderData)

	resDom.Data = new(domSchema.TotalCountryData)

	resDom.Data.CountryInfo = new(domSchema.CountryInfo)
	resDom.Data.CountryInfo.Code = res.Result.PageContext.CountryCode
	resDom.Data.CountryInfo.Name = res.Result.PageContext.Feature.Properties.Adm0Name
	resDom.Data.CountryInfo.ISO2Code = res.Result.PageContext.Feature.Properties.ISO2Code
	resDom.Data.CountryInfo.ISO3Code = res.Result.PageContext.Feature.Properties.ISO3Code
	resDom.Data.CountryInfo.WHORegion = res.Result.PageContext.Feature.Properties.WHORegion

	resDom.Data.StartDate = res.Result.PageContext.StartDate
	resDom.Data.LastUpdate = res.Result.PageContext.LastUpdate

	resDom.Data.TodayData = new(domSchema.DayData)
	resDom.Data.TodayData.Confirmed = res.Result.PageContext.Today.Confirmed
	resDom.Data.TodayData.Death = res.Result.PageContext.Today.Deaths
	resDom.Data.TodayData.CumulativeConfirmed = res.Result.PageContext.Today.CumulativeConfirmed
	resDom.Data.TodayData.CumulativeDeath = res.Result.PageContext.Today.CumulativeDeaths

	resDom.Data.YesterdayData = new(domSchema.DayData)
	resDom.Data.YesterdayData.Confirmed = res.Result.PageContext.Yesterday.Confirmed
	resDom.Data.YesterdayData.Death = res.Result.PageContext.Yesterday.Deaths
	resDom.Data.YesterdayData.CumulativeConfirmed = res.Result.PageContext.Yesterday.CumulativeConfirmed
	resDom.Data.YesterdayData.CumulativeDeath = res.Result.PageContext.Yesterday.CumulativeDeaths

	resDom.Data.TotalData = new(domSchema.TotalData)
	resDom.Data.TotalData.Confirmed = res.Result.PageContext.Totals.Confirmed
	resDom.Data.TotalData.Death = res.Result.PageContext.Totals.Deaths
	resDom.Data.TotalData.RegionCumulativeConfirmed = res.Result.PageContext.Totals.CumulativeConfirmed
	resDom.Data.TotalData.RegionCumulativeDeath = res.Result.PageContext.Totals.CumulativeDeaths

	return resDom, nil
}
