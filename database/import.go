package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/Masterminds/squirrel"
)

func ImportData(db *sql.DB, documentData io.Reader, segmentData io.Reader) error {
	if err := importDocumentData(db, documentData); err != nil {
		return err
	}

	if err := importSegmentData(db, segmentData); err != nil {
		return err
	}

	return nil
}

func importDocumentData(db *sql.DB, reader io.Reader) error {
	csvReader := csv.NewReader(reader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv document data: %w", err)
	}

	builder := squirrel.Insert(TableDocumentKop).Columns(
		ColumnBedrijfsNummer,
		ColumnDocumentNummer,
		ColumnBoekJaar,
		ColumnDocumentSoort,
		ColumnDocumentDatum,
		ColumnBoekingDatum,
		ColumnBoekMaand,
		ColumnInvoerDatum,
		ColumnInvoerTijd,
	)

	for _, row := range rows {
		// Convert []string to []any
		values := make([]any, len(row))
		for index, value := range row {
			values[index] = value
		}

		// Append values to query
		builder = builder.Values(values...)
	}

	query, arguments, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to construct query: %w", err)
	}

	if _, err := db.Exec(query, arguments...); err != nil {
		return fmt.Errorf("failed to insert document data: %w", err)
	}

	return nil
}

func importSegmentData(db *sql.DB, reader io.Reader) error {
	csvReader := csv.NewReader(reader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv segment data: %w", err)
	}

	builder := squirrel.Insert(TableDocumentSegment).Columns(
		ColumnBedrijfsNummer,
		ColumnDocumentNummer,
		ColumnBoekJaar,
		ColumnBoekingsregelNummer,
		ColumnBoekingRegelID,
		ColumnVereffeningDatum,
		ColumnVereffeningInvoerDatum,
		ColumnVereffeningsDocumentNummer,
		ColumnBoekingssleutel,
	)

	for _, row := range rows {
		// Convert []string to []any
		values := make([]any, len(row))
		for index, value := range row {
			values[index] = value
		}

		// Append values to query
		builder = builder.Values(values...)
	}

	query, arguments, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to construct query: %w", err)
	}

	if _, err := db.Exec(query, arguments...); err != nil {
		return fmt.Errorf("failed to insert segment data: %w", err)
	}

	return nil
}
