package model

type DocumentSegment struct {
	Bedrijfsnummer            Bedrijfsnummer
	DocumentNummer            DocumentNummer
	BoekJaar                  Date
	Regelnummer               int
	RegelIdentificatie        string
	VereffeningDatum          Date
	VereffeningInvoerDatum    Date
	VereffeningDocumentNummer DocumentNummer
	BoekingSleutel            BoekingSleutel
}
