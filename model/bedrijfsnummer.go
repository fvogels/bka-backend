package model

import (
	"fmt"
	"unicode/utf8"
)

type Bedrijfsnummer struct {
	value string
}

func NewBedrijfsnummer(str string) Bedrijfsnummer {
	if utf8.RuneCountInString(str) != 4 {
		panic(fmt.Sprintf("invalid bedrijfsnummer %s", str))
	}

	return Bedrijfsnummer{
		value: str,
	}
}

func (nummer Bedrijfsnummer) String() string {
	return nummer.value
}

func (nummer Bedrijfsnummer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, nummer.value)), nil
}
