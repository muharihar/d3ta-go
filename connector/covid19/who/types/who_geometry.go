package types

type WHOCountryByGeometry struct {
	Type    string  `json:"type"`
	Objects Objects `json:"objects"`
}

type Objects struct {
	Countries Countries `json:"countries"`
}

type Countries struct {
	Type       string            `json:"type"`
	Geometries []GeometryElement `json:"geometries"`
}

type GeometryElement struct {
	Arcs       interface{}        `json:"arcs"`
	Type       *Type              `json:"type"`
	Properties GeometryProperties `json:"properties"`
}

type GeometryProperties struct {
	ID        interface{} `json:"id"`
	WHORegion WHORegion   `json:"WHO_REGION"`
	ISO2Code  string      `json:"ISO_2_CODE"`
	Adm0Name  string      `json:"ADM0_NAME"`
	ISO3Code  string      `json:"ISO_3_CODE"`
}

type WHORegion string

const (
	Afro  WHORegion = "AFRO"
	Amro  WHORegion = "AMRO"
	Emro  WHORegion = "EMRO"
	Euro  WHORegion = "EURO"
	Searo WHORegion = "SEARO"
	Wpro  WHORegion = "WPRO"
)

type Type string

const (
	MultiPolygon Type = "MultiPolygon"
	Polygon      Type = "Polygon"
)
