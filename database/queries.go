package database

import (
	"bass-backend/database/filters"
	"bass-backend/database/names"
	"bass-backend/model"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func CountDocuments(db *sql.DB, filter filters.Filter) (int, error) {
	query, arguments, err := squirrel.Select("COUNT(*)").From(names.TableDocumentKop).Where(filter.Build()).ToSql()

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
	query, arguments, err := squirrel.Insert(names.TableDocumentKop).Columns(
		names.ColumnBedrijfsNummer,
		names.ColumnDocumentNummer,
		names.ColumnBoekJaar,
		names.ColumnDocumentSoort,
		names.ColumnDocumentDatum,
		names.ColumnBoekingDatum,
		names.ColumnBoekMaand,
		names.ColumnInvoerDatum,
		names.ColumnInvoerTijd,
	).Values(
		kop.Bedrijfsnummer.String(),
		kop.DocumentNummer.String(),
		kop.BoekJaar.String(),
		kop.DocumentSoort.String(),
		kop.DocumentDatum.ToYYYYMMSS(),
		kop.BoekingsDatum.ToYYYYMMSS(),
		kop.Boekmaand.String(),
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
	query, arguments, err := squirrel.Insert(names.TableDocumentSegment).Columns(
		names.ColumnBedrijfsNummer,
		names.ColumnDocumentNummer,
		names.ColumnBoekJaar,
		names.ColumnBoekingRegelID,
		names.ColumnVereffeningDatum,
		names.ColumnVereffeningInvoerDatum,
		names.ColumnVereffeningsDocumentNummer,
		names.ColumnBoekingssleutel,
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
