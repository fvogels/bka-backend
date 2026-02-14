package model

import (
	"fmt"
	"unicode/utf8"
)

type Bedrijfsnummer struct {
	value string
}

func NewBedrijfsnummer(str string) Bedrijfsnummer {
	result, err := ParseBedrijfsnummer(str)

	if err != nil {
		panic(fmt.Sprintf("invalid bedrijfsnummer %s", str))
	}

	return result
}

func ParseBedrijfsnummer(str string) (Bedrijfsnummer, error) {
	if utf8.RuneCountInString(str) != 4 {
		return Bedrijfsnummer{}, fmt.Errorf("%w: %s", ErrInvalidString, str)
	}

	return Bedrijfsnummer{
		value: str,
	}, nil
}

func (nummer Bedrijfsnummer) String() string {
	return nummer.value
}

func (nummer Bedrijfsnummer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, nummer.value)), nil
}
