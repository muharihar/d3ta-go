package adapter

import domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"

// ICovid19Adapter interface
type ICovid19Adapter interface {
	DisplayCurrentDataByCountry(req *domSchema.DisplayCurrentDataByCountryRequest) (*domSchema.TotalCountryProviderData, error)
}

// BaseCovid19Adapter represent BaseCovid19Adapter
type BaseCovid19Adapter struct {
	info Covid19AdapterInfo
}

// SetInfo set Info
func (b *BaseCovid19Adapter) SetInfo(info Covid19AdapterInfo) {
	b.info = info
}

// GetInfo set Info
func (b *BaseCovid19Adapter) GetInfo() Covid19AdapterInfo {
	return b.info
}
