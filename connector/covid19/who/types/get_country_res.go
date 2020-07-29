package types

type GetCountryResponse struct {
	ComponentChunkName string        `json:"componentChunkName"`
	Path               string        `json:"path"`
	Result             CountryResult `json:"result"`
}

type CountryResult struct {
	PageContext CountryPageContext `json:"pageContext"`
}

type CountryPageContext struct {
	CountryCode          string      `json:"countryCode"`
	HideMap              bool        `json:"hideMap"`
	ByCountry            interface{} `json:"byCountry"`
	CountriesDailyChange interface{} `json:"countriesDailyChange"`
	CountryGroup         interface{} `json:"countryGroup"`
	StartDate            string      `json:"startDate"`
	Today                Day         `json:"today"`
	Yesterday            Day         `json:"yesterday"` // Yesterday   map[string]int64 `json:"yesterday"`
	Totals               Totals      `json:"totals"`    // Totals  map[string]int64 `json:"totals"`
	CountryVals          interface{} `json:"countryVals"`
	Feature              Feature     `json:"feature"`
	ByDayCumulative      interface{} `json:"byDayCumulative"`
	LastUpdate           string      `json:"lastUpdate"`
	ByDay                interface{} `json:"byDay"`
	TransmissionData     interface{} `json:"transmissionData"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	ISO2Code  string `json:"ISO_2_CODE"`
	Adm0Name  string `json:"ADM0_NAME"`
	Adm0VizN  string `json:"ADM0_VIZ_N"`
	ISO3Code  string `json:"ISO_3_CODE"`
	WHORegion string `json:"WHO_REGION"`
}

type Day struct {
	Confirmed           int64 `json:"Confirmed"`
	CumulativeConfirmed int64 `json:"Cumulative Confirmed"`
	Deaths              int64 `json:"Deaths"`
	CumulativeDeaths    int64 `json:"Cumulative Deaths"`
}

type Totals struct {
	Deaths                int64 `json:"Deaths"`
	CumulativeDeaths      int64 `json:"Cumulative Deaths"`
	DeathsLast7Days       int64 `json:"Deaths Last 7 Days"`
	DeathsLast7DaysChange int64 `json:"Deaths Last 7 Days Change"`
	DeathsPerMillion      int64 `json:"Deaths Per Million"`
	Confirmed             int64 `json:"Confirmed"`
	CumulativeConfirmed   int64 `json:"Cumulative Confirmed"`
	CasesLast7Days        int64 `json:"Cases Last 7 Days"`
	CasesLast7DaysChange  int64 `json:"Cases Last 7 Days Change"`
	CasesPerMillion       int64 `json:"Cases Per Million"`
}
