package meta

var (
	DocumentKop = DocumentKopMetadata{
		Table:          "BKPF",
		Bedrijfsnummer: "BUKRS",
		Documentnummer: "BELNR",
		Boekjaar:       "GJAHR",
		Documentsoort:  "BLART",
		Documentdatum:  "BLDAT",
		Boekingdatum:   "BUDAT",
		Boekmaand:      "MONAT",
		Invoerdatum:    "CPUDT",
		Invoertijd:     "CPUTM",
	}

	DocumentSegment = DocumentSegmentMetadata{
		Table:                      "BSEG",
		Bedrijfsnummer:             "BUKRS",
		Documentnummer:             "BELNR",
		Boekjaar:                   "GJAHR",
		Boekingsregelnummer:        "BUZEI",
		BoekingregelID:             "BUZID",
		Vereffeningdatum:           "AUGDT",
		Vereffeninginvoerdatum:     "AUGCP",
		Vereffeningsdocumentnummer: "AUGBL",
		Boekingssleutel:            "BSCHL",
	}
)

type DocumentKopMetadata struct {
	Table          string
	Bedrijfsnummer string
	Documentnummer string
	Boekjaar       string
	Documentsoort  string
	Documentdatum  string
	Boekingdatum   string
	Boekmaand      string
	Invoerdatum    string
	Invoertijd     string
}

type DocumentSegmentMetadata struct {
	Table                      string
	Bedrijfsnummer             string
	Documentnummer             string
	Boekjaar                   string
	Boekingsregelnummer        string
	BoekingregelID             string
	Vereffeningdatum           string
	Vereffeninginvoerdatum     string
	Vereffeningsdocumentnummer string
	Boekingssleutel            string
}
