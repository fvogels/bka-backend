package model

import "fmt"

type DocumentNummer struct {
	value string
}

func NewDocumentNummer(documentNummer string) DocumentNummer {
	if len(documentNummer) != 10 {
		panic(fmt.Sprintf("invalid documentnummer %d", documentNummer))
	}

	return DocumentNummer{
		value: documentNummer,
	}
}

func (nummer DocumentNummer) String() string {
	return nummer.value
}

func (nummer DocumentNummer) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, nummer.value), nil
}
