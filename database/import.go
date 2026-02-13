package database

import (
	"bass-backend/database/names"
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

	builder := squirrel.Insert(names.TableDocumentKop).Columns(
		names.ColumnBedrijfsNummer,
		names.ColumnDocumentNummer,
		names.ColumnBoekJaar,
		names.ColumnDocumentSoort,
		names.ColumnDocumentDatum,
		names.ColumnBoekingDatum,
		names.ColumnBoekMaand,
		names.ColumnInvoerDatum,
		names.ColumnInvoerTijd,
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

	builder := squirrel.Insert(names.TableDocumentSegment).Columns(
		names.ColumnBedrijfsNummer,
		names.ColumnDocumentNummer,
		names.ColumnBoekJaar,
		names.ColumnBoekingsregelNummer,
		names.ColumnBoekingRegelID,
		names.ColumnVereffeningDatum,
		names.ColumnVereffeningInvoerDatum,
		names.ColumnVereffeningsDocumentNummer,
		names.ColumnBoekingssleutel,
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
