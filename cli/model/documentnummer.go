package model

import "fmt"

type DocumentNummer struct {
	value int
}

func NewDocumentNummer(n int) DocumentNummer {
	return DocumentNummer{
		value: n,
	}
}

func (nummer DocumentNummer) String() string {
	return fmt.Sprintf("%010d", nummer.value)
}
