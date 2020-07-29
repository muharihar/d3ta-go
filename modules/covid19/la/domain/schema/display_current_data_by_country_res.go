package schema

import "encoding/json"

// DisplayCurrentDataByCountryResponse represent DisplayCurrentDataByCountryResponse
type DisplayCurrentDataByCountryResponse struct {
	Query interface{}                   `json:"query"`
	Data  *TotalCountryProviderDataList `json:"data"`
}

// ToJSON covert to JSON
func (r *DisplayCurrentDataByCountryResponse) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

// TotalCountryProviderDataList represent TotalCountryProviderDataList
type TotalCountryProviderDataList []*TotalCountryProviderData

// TotalCountryProviderData represent TotalCountryProviderData
type TotalCountryProviderData struct {
	Provider    string            `json:"provider"`
	Information string            `json:"information"`
	Data        *TotalCountryData `json:"data"`
}

// TotalCountryData represent Total Country Data
type TotalCountryData struct {
	CountryInfo   *CountryInfo `json:"countryInfo"`
	StartDate     string       `json:"startDate"`
	LastUpdate    string       `json:"lastUpdate"`
	TodayData     *DayData     `json:"today"`
	YesterdayData *DayData     `json:"yesterday"`
	TotalData     *TotalData   `json:"total"`
}

// CountryInfo represent Country Information
type CountryInfo struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	ISO2Code  string `json:"ISO2Code"`
	ISO3Code  string `json:"ISO3Code"`
	WHORegion string `json:"WHORegion"`
}

// DayData represent Day Data
type DayData struct {
	Death               int64 `json:"death"`
	Confirmed           int64 `json:"confirmed"`
	CumulativeDeath     int64 `json:"cumulativeDeath"`
	CumulativeConfirmed int64 `json:"cumulativeConfirmed"`
}

// TotalData represent Total Data
type TotalData struct {
	Death                     int64 `json:"death"`
	Confirmed                 int64 `json:"confirmed"`
	RegionCumulativeDeath     int64 `json:"regionCumulativeDeath"`
	RegionCumulativeConfirmed int64 `json:"regionCumulativeConfirmed"`
}
