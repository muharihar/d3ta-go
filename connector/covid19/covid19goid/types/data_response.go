package types

type DataResponse struct {
	LastUpdate string    `json:"last_update"`
	Kasus      Kasus     `json:"kasus"`
	Sembuh     Meninggal `json:"sembuh"`
	Meninggal  Meninggal `json:"meninggal"`
	Perawatan  Meninggal `json:"perawatan"`
}

type Kasus struct {
	KondisiPenyerta GejalaClass       `json:"kondisi_penyerta"`
	JenisKelamin    GejalaClass       `json:"jenis_kelamin"`
	KelompokUmur    KasusKelompokUmur `json:"kelompok_umur"`
	Gejala          GejalaClass       `json:"gejala"`
}

type GejalaClass struct {
	CurrentData int64             `json:"current_data"`
	MissingData float64           `json:"missing_data"`
	ListData    []GejalaListDatum `json:"list_data"`
}

type GejalaListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
}

type KasusKelompokUmur struct {
	CurrentData int64                   `json:"current_data"`
	MissingData float64                 `json:"missing_data"`
	ListData    []KelompokUmurListDatum `json:"list_data"`
}

type KelompokUmurListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
	Usia     Usia    `json:"usia"`
}

type Meninggal struct {
	KondisiPenyerta PurpleGejala          `json:"kondisi_penyerta"`
	JenisKelamin    PurpleGejala          `json:"jenis_kelamin"`
	KelompokUmur    MeninggalKelompokUmur `json:"kelompok_umur"`
	Gejala          PurpleGejala          `json:"gejala"`
}

type PurpleGejala struct {
	ListData []GejalaListDatum `json:"list_data"`
}

type MeninggalKelompokUmur struct {
	ListData []KelompokUmurListDatum `json:"list_data"`
}
