package database

import (
	"bass-backend/database/meta"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"

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

	builder := squirrel.Insert(meta.DocumentKop.Table).Columns(
		meta.DocumentKop.Bedrijfsnummer,
		meta.DocumentKop.Documentnummer,
		meta.DocumentKop.Boekjaar,
		meta.DocumentKop.Documentsoort,
		meta.DocumentKop.Documentdatum,
		meta.DocumentKop.Boekingdatum,
		meta.DocumentKop.Boekmaand,
		meta.DocumentKop.Invoerdatum,
		meta.DocumentKop.Invoertijd,
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

	block := [][]string{}
	for index, row := range rows {
		block = append(block, row)

		if index%100 == 0 {
			importSegmentDataBlock(db, block)
			block = [][]string{}
		}
	}

	importSegmentDataBlock(db, block)

	return nil
}

func importSegmentDataBlock(db *sql.DB, rows [][]string) error {
	slog.Debug("Writing block of segments", slog.Int("rowCount", len(rows)))

	builder := squirrel.Insert(meta.DocumentSegment.Table).Columns(
		meta.DocumentSegment.Bedrijfsnummer,
		meta.DocumentSegment.Documentnummer,
		meta.DocumentSegment.Boekjaar,
		meta.DocumentSegment.Boekingsregelnummer,
		meta.DocumentSegment.BoekingregelID,
		meta.DocumentSegment.Vereffeningdatum,
		meta.DocumentSegment.Vereffeninginvoerdatum,
		meta.DocumentSegment.Vereffeningsdocumentnummer,
		meta.DocumentSegment.Boekingssleutel,
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
