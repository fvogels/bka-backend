package model

type DocumentKop struct {
	Bedrijfsnummer Bedrijfsnummer
	DocumentNummer DocumentNummer
	BoekJaar       Date
	DocumentSoort  DocumentSoort
	DocumentDatum  Date
	BoekingsDatum  Date
	Boekmaand      int
	InvoerDatum    Date
	InvoerTijd     Time
}
