package model

import "unicode/utf8"

type Bedrijfsnummer struct {
	value string
}

func NewBedrijfsnummer(str string) Bedrijfsnummer {
	if utf8.RuneCountInString(str) != 4 {
		panic("invalid bedrijfsnummer")
	}

	return Bedrijfsnummer{
		value: str,
	}
}

func (nummer Bedrijfsnummer) String() string {
	return nummer.value
}
