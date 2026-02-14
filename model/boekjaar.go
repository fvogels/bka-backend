package model

import "fmt"

type BoekJaar string

func NewBoekJaar(str string) BoekJaar {
	result, err := ParseBoekJaar(str)

	if err != nil {
		panic(fmt.Sprintf("Invalid boekjaar: %s", err.Error()))
	}

	return result
}

func ParseBoekJaar(str string) (BoekJaar, error) {
	if len(str) != 4 {
		return "", fmt.Errorf("%w: %s", ErrInvalidString, str)
	}

	return BoekJaar(str), nil
}

func (boekmaand BoekJaar) String() string {
	return string(boekmaand)
}
