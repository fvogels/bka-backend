package model

type DocumentSegment struct {
	Bedrijfsnummer            Bedrijfsnummer
	DocumentNummer            DocumentNummer
	BoekJaar                  Date
	Regelnummer               int
	RegelIdentificatie        rune
	VereffeningDatum          Date
	VereffeningInvoerDatum    Date
	VereffeningDocumentNummer DocumentNummer
	BoekingSleutel            BoekingSleutel
}
