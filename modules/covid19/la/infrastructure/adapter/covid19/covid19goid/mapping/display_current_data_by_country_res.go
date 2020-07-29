package mapping

import (
	con19type "github.com/muharihar/d3ta-go/connector/covid19/covid19goid/types"
	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
)

// MapDisplayCurrentDataByCountryRes mapping DisplayCurrentDataByCountryRes
func MapDisplayCurrentDataByCountryRes(res *con19type.UpdateResponse) (*domSchema.TotalCountryProviderData, error) {

	resDom := new(domSchema.TotalCountryProviderData)

	resDom.Data = new(domSchema.TotalCountryData)

	resDom.Data.CountryInfo = new(domSchema.CountryInfo)
	resDom.Data.CountryInfo.Code = "ID"
	resDom.Data.CountryInfo.Name = "INDONESIA"
	resDom.Data.CountryInfo.ISO2Code = "ID"
	resDom.Data.CountryInfo.ISO3Code = "IDN"
	resDom.Data.CountryInfo.WHORegion = "SEARO"

	resDom.Data.StartDate = res.Update.Harian[0].KeyAsString
	resDom.Data.LastUpdate = res.Update.Harian[len(res.Update.Harian)-1].KeyAsString

	resDom.Data.TodayData = new(domSchema.DayData)
	resDom.Data.TodayData.Confirmed = res.Update.Harian[len(res.Update.Harian)-1].JumlahPositif.Value
	resDom.Data.TodayData.Death = res.Update.Harian[len(res.Update.Harian)-1].JumlahMeninggal.Value
	resDom.Data.TodayData.CumulativeConfirmed = res.Update.Harian[len(res.Update.Harian)-1].JumlahPositifKum.Value
	resDom.Data.TodayData.CumulativeDeath = res.Update.Harian[len(res.Update.Harian)-1].JumlahMeninggalKum.Value

	resDom.Data.YesterdayData = new(domSchema.DayData)
	resDom.Data.YesterdayData.Confirmed = res.Update.Harian[len(res.Update.Harian)-2].JumlahPositif.Value
	resDom.Data.YesterdayData.Death = res.Update.Harian[len(res.Update.Harian)-2].JumlahMeninggal.Value
	resDom.Data.YesterdayData.CumulativeConfirmed = res.Update.Harian[len(res.Update.Harian)-2].JumlahPositifKum.Value
	resDom.Data.YesterdayData.CumulativeDeath = res.Update.Harian[len(res.Update.Harian)-2].JumlahMeninggalKum.Value

	resDom.Data.TotalData = new(domSchema.TotalData)
	resDom.Data.TotalData.Confirmed = res.Update.Total.JumlahPositif
	resDom.Data.TotalData.Death = res.Update.Total.JumlahMeninggal
	resDom.Data.TotalData.RegionCumulativeConfirmed = 0
	resDom.Data.TotalData.RegionCumulativeDeath = 0

	return resDom, nil
}
