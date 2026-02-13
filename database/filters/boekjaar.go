package filters

import (
	"bass-backend/database/names"

	"github.com/Masterminds/squirrel"
)

func Boekjaar(boekjaar int) Filter {
	return boekjaarFilter{boekjaar: boekjaar}
}

type boekjaarFilter struct {
	boekjaar int
}

func (filter boekjaarFilter) Build() squirrel.Sqlizer {
	return squirrel.Eq{names.ColumnBoekJaar: filter.boekjaar}
}
