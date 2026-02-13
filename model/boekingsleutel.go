package model

type BoekingSleutel string

func (sleutel BoekingSleutel) String() string {
	return string(sleutel)
}

const (
	Debit  BoekingSleutel = "40"
	Credit BoekingSleutel = "50"
)
