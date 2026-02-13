package model

type BoekMaand string

func NewBoekMaand(str string) BoekJaar {
	if len(str) != 2 {
		panic("invalid boekmaand")
	}

	return BoekJaar(str)
}

func (boekmaand BoekMaand) String() string {
	return string(boekmaand)
}
