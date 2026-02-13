package filters

import (
	"bass-backend/database/names"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func Bedrijfsnummer(bedrijfsnummer string) Filter {
	return bedrijfsnummerFilter{
		bedrijfsnummer: bedrijfsnummer,
	}
}

type bedrijfsnummerFilter struct {
	bedrijfsnummer string
}

func (filter bedrijfsnummerFilter) Build() squirrel.Sqlizer {
	return squirrel.Eq{fmt.Sprintf("%s.%s", names.TableDocumentKop, names.ColumnBedrijfsNummer): filter.bedrijfsnummer}
}
