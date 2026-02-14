package database

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func InsertDocumentKop(db *sql.DB, kop model.DocumentKop) error {
	query, arguments, err := squirrel.Insert(meta.DocumentKop.Table).Columns(
		meta.DocumentKop.Bedrijfsnummer,
		meta.DocumentKop.Documentnummer,
		meta.DocumentKop.Boekjaar,
		meta.DocumentKop.Documentsoort,
		meta.DocumentKop.Documentdatum,
		meta.DocumentKop.Boekingdatum,
		meta.DocumentKop.Boekmaand,
		meta.DocumentKop.Invoerdatum,
		meta.DocumentKop.Invoertijd,
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
	query, arguments, err := squirrel.Insert(meta.DocumentSegment.Table).Columns(
		meta.DocumentSegment.Bedrijfsnummer,
		meta.DocumentSegment.Documentnummer,
		meta.DocumentSegment.Boekjaar,
		meta.DocumentSegment.BoekingregelID,
		meta.DocumentSegment.Vereffeningdatum,
		meta.DocumentSegment.Vereffeninginvoerdatum,
		meta.DocumentSegment.Vereffeningsdocumentnummer,
		meta.DocumentSegment.Boekingssleutel,
	).Values(
		segment.Bedrijfsnummer.String(),
		segment.DocumentNummer.String(),
		segment.BoekJaar.ToYYYYMMSS(),
		fmt.Sprintf("%03d", segment.Regelnummer),
		segment.RegelIdentificatie,
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
