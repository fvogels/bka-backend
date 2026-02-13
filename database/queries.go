package database

import (
	"bass-backend/cli/model"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type SearchPredicate interface {
	Build() any
}

type BoekjaarSearchPredicate struct {
	Boekjaar model.Date
}

func (predicate BoekjaarSearchPredicate) Build() any {
	return squirrel.Eq{ColumnBoekJaar: predicate.Boekjaar.ToYYYYMMSS()}
}

type BedrijfsnummerSearchPredicate struct {
	Bedrijfsnummer string
}

func CountDocuments(db *sql.DB, predicate SearchPredicate) (int, error) {
	query, arguments, err := squirrel.Select("COUNT(*)").From(TableDocumentKop).Where(predicate.Build()).ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to construct SQL query: %w", err)
	}

	row := db.QueryRow(query, arguments...)

	var documentCount int
	if err := row.Scan(&documentCount); err != nil {
		return 0, fmt.Errorf("failed to count the number of documents: %w", err)
	}

	return documentCount, nil
}

func InsertDocumentKop(db *sql.DB, kop model.DocumentKop) error {
	query, arguments, err := squirrel.Insert(TableDocumentKop).Columns(
		ColumnBedrijfsNummer,
		ColumnDocumentNummer,
		ColumnBoekJaar,
		ColumnDocumentSoort,
		ColumnDocumentDatum,
		ColumnBoekingDatum,
		ColumnBoekMaand,
		ColumnInvoerDatum,
		ColumnInvoerTijd,
	).Values(
		kop.Bedrijfsnummer.String(),
		kop.DocumentNummer.String(),
		kop.BoekJaar.ToYYYYMMSS(),
		kop.DocumentSoort.String(),
		kop.DocumentDatum.ToYYYYMMSS(),
		kop.BoekingsDatum.ToYYYYMMSS(),
		fmt.Sprintf("%02d", kop.Boekmaand),
		kop.InvoerDatum.ToYYYYMMSS(),
		kop.InvoerTijd.ToHHMMSS(),
	).ToSql()

	if err != nil {
		return fmt.Errorf("failed to create query for documentkop insertion: %w", err)
	}

	if _, err := db.Exec(query, arguments...); err != nil {
		return fmt.Errorf("failed to insert documentkop: %w", err)
	}

	return nil
}

func InsertDocumentSegment(db *sql.DB, segment model.DocumentSegment) error {
	query, arguments, err := squirrel.Insert(TableDocumentSegment).Columns(
		ColumnBedrijfsNummer,
		ColumnDocumentNummer,
		ColumnBoekJaar,
		ColumnBoekingRegelID,
		ColumnVereffeningDatum,
		ColumnVereffeningInvoerDatum,
		ColumnVereffeningsDocumentNummer,
		ColumnBoekingssleutel,
	).Values(
		segment.Bedrijfsnummer.String(),
		segment.DocumentNummer.String(),
		segment.BoekJaar.ToYYYYMMSS(),
		fmt.Sprintf("%03d", segment.Regelnummer),
		string([]rune{segment.RegelIdentificatie}),
		segment.VereffeningDatum.ToYYYYMMSS(),
		segment.VereffeningInvoerDatum.ToYYYYMMSS(),
		segment.VereffeningDocumentNummer.String(),
		segment.BoekingSleutel.String(),
	).ToSql()

	if err != nil {
		return fmt.Errorf("failed to create query for documentsegment insertion: %w", err)
	}

	if _, err := db.Exec(query, arguments...); err != nil {
		return fmt.Errorf("failed to insert documentsegment: %w", err)
	}

	return nil
}
