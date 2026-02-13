package model

type BoekJaar string

func NewBoekJaar(str string) BoekJaar {
	if len(str) != 4 {
		panic("invalid boekjaar")
	}

	return BoekJaar(str)
}

func (boekmaand BoekJaar) String() string {
	return string(boekmaand)
}
