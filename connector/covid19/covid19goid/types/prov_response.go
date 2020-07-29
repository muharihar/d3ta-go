package types

type ProvResponse struct {
	LastDate      string      `json:"last_date"`
	CurrentData   float64     `json:"current_data"`
	MissingData   float64     `json:"missing_data"`
	TanpaProvinsi int64       `json:"tanpa_provinsi"`
	ListData      []ListDatum `json:"list_data"`
}

type ListDatum struct {
	Key             string         `json:"key"`
	DocCount        float64        `json:"doc_count"`
	JumlahKasus     int64          `json:"jumlah_kasus"`
	JumlahSembuh    int64          `json:"jumlah_sembuh"`
	JumlahMeninggal int64          `json:"jumlah_meninggal"`
	JumlahDirawat   int64          `json:"jumlah_dirawat"`
	JenisKelamin    []JenisKelamin `json:"jenis_kelamin"`
	KelompokUmur    []KelompokUmur `json:"kelompok_umur"`
	Lokasi          Lokasi         `json:"lokasi"`
	Penambahan      PenambahanProv `json:"penambahan"`
}

type JenisKelamin struct {
	Key      JenisKelaminKey `json:"key"`
	DocCount int64           `json:"doc_count"`
}

type KelompokUmur struct {
	Key      KelompokUmurKey `json:"key"`
	DocCount int64           `json:"doc_count"`
	Usia     Usia            `json:"usia"`
}

type Lokasi struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type PenambahanProv struct {
	Positif   int64 `json:"positif"`
	Sembuh    int64 `json:"sembuh"`
	Meninggal int64 `json:"meninggal"`
}

type JenisKelaminKey string

const (
	LakiLaki  JenisKelaminKey = "LAKI-LAKI"
	Perempuan JenisKelaminKey = "PEREMPUAN"
)

type KelompokUmurKey string

const (
	The05   KelompokUmurKey = "0-5"
	The1830 KelompokUmurKey = "18-30"
	The3145 KelompokUmurKey = "31-45"
	The4659 KelompokUmurKey = "46-59"
	The60   KelompokUmurKey = "≥ 60"
	The617  KelompokUmurKey = "6-17"
)
