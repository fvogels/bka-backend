package model

type DocumentNummer struct {
	value string
}

func NewDocumentNummer(documentNummer string) DocumentNummer {
	if len(documentNummer) != 10 {
		panic("invalid documentnummer")
	}

	return DocumentNummer{
		value: documentNummer,
	}
}

func (nummer DocumentNummer) String() string {
	return nummer.value
}
