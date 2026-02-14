package meta

var (
	DocumentKop = DocumentKopMetadata{
		Table:          "BKPF",
		BedrijfsNummer: "BUKRS",
		DocumentNummer: "BELNR",
		BoekJaar:       "GJAHR",
		DocumentSoort:  "BLART",
		DocumentDatum:  "BLDAT",
		BoekingDatum:   "BUDAT",
		BoekMaand:      "MONAT",
		InvoerDatum:    "CPUDT",
		InvoerTijd:     "CPUTM",
	}

	DocumentSegment = DocumentSegmentMetadata{
		Table:                      "BSEG",
		BedrijfsNummer:             "BUKRS",
		DocumentNummer:             "BELNR",
		BoekJaar:                   "GJAHR",
		BoekingsregelNummer:        "BUZEI",
		BoekingRegelID:             "BUZID",
		VereffeningDatum:           "AUGDT",
		VereffeningInvoerDatum:     "AUGCP",
		VereffeningsDocumentNummer: "AUGBL",
		Boekingssleutel:            "BSCHL",
	}
)

type DocumentKopMetadata struct {
	Table          string
	BedrijfsNummer string
	DocumentNummer string
	BoekJaar       string
	DocumentSoort  string
	DocumentDatum  string
	BoekingDatum   string
	BoekMaand      string
	InvoerDatum    string
	InvoerTijd     string
}

type DocumentSegmentMetadata struct {
	Table                      string
	BedrijfsNummer             string
	DocumentNummer             string
	BoekJaar                   string
	BoekingsregelNummer        string
	BoekingRegelID             string
	VereffeningDatum           string
	VereffeningInvoerDatum     string
	VereffeningsDocumentNummer string
	Boekingssleutel            string
}
