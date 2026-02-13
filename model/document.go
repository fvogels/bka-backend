package model

type Document struct {
	Bedrijfsnummer Bedrijfsnummer `json:"bedrijfsnummer"`
	DocumentNummer DocumentNummer `json:"documentnummer"`
	BoekJaar       BoekJaar       `json:"boekjaar"`
	DocumentSoort  DocumentSoort  `json:"soort"`
	DocumentDatum  Date           `json:"documentdatum"`
	BoekingsDatum  Date           `json:"boekingsdatum"`
	Boekmaand      BoekMaand      `json:"boekmaand"`
	InvoerDatum    Date           `json:"invoerdatum"`
	InvoerTijd     Time           `json:"invoertijd"`
	Segmenten      []Segment      `json:"segmenten"`
}

type Segment struct {
	Regelnummer               int            `json:"regelnummer"`
	RegelIdentificatie        string         `json:"regelidentificatie"`
	VereffeningDatum          Date           `json:"vereffeningsdatum"`
	VereffeningInvoerDatum    Date           `json:"vereffeningsinvoerdatum"`
	VereffeningDocumentNummer DocumentNummer `json:"vereffeningsdocumentnummer"`
	BoekingSleutel            BoekingSleutel `json:"boekingssleutel"`
}
