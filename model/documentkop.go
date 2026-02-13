package model

type DocumentKop struct {
	Bedrijfsnummer Bedrijfsnummer
	DocumentNummer DocumentNummer
	BoekJaar       BoekJaar
	DocumentSoort  DocumentSoort
	DocumentDatum  Date
	BoekingsDatum  Date
	Boekmaand      BoekMaand
	InvoerDatum    Date
	InvoerTijd     Time
}
