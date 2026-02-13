package database

import (
	"bass-backend/database/filters"
	"bass-backend/database/names"
	"bass-backend/model"
	"bass-backend/util"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/hoisie/mustache"
)

func ListDocuments(db *sql.DB, filter filters.Filter) ([]*model.Document, error) {
	nameTable := map[string]string{
		"koptabel":                   names.TableDocumentKop,
		"segmenttabel":               names.TableDocumentSegment,
		"bedrijfsnummer":             names.ColumnBedrijfsNummer,
		"documentnummer":             names.ColumnDocumentNummer,
		"boekjaar":                   names.ColumnBoekJaar,
		"documentsoort":              names.ColumnDocumentSoort,
		"documentdatum":              names.ColumnDocumentDatum,
		"boekingsdatum":              names.ColumnBoekingDatum,
		"boekmaand":                  names.ColumnBoekMaand,
		"invoerdatum":                names.ColumnInvoerDatum,
		"invoertijd":                 names.ColumnInvoerTijd,
		"boekingsregelnummer":        names.ColumnBoekingsregelNummer,
		"boekingsregelidentificatie": names.ColumnBoekingRegelID,
		"vereffeningsdatum":          names.ColumnVereffeningDatum,
		"vereffeningsinvoerdatum":    names.ColumnVereffeningInvoerDatum,
		"vereffeningsdocument":       names.ColumnVereffeningsDocumentNummer,
		"boekingssleutel":            names.ColumnBoekingssleutel,
	}

	query, arguments, err := squirrel.Select(
		mustache.Render(`
			{{koptabel}}.{{bedrijfsnummer}},
			{{koptabel}}.{{documentnummer}},
			{{koptabel}}.{{boekjaar}},
			{{documentsoort}},
			{{documentdatum}},
			{{boekingsdatum}},
			{{boekmaand}},
			{{invoerdatum}},
			{{invoertijd}},
			{{boekingsregelnummer}},
			{{boekingsregelidentificatie}},
			{{vereffeningsdatum}},
			{{vereffeningsinvoerdatum}},
			{{vereffeningsdocument}},
			{{boekingssleutel}}
		`, nameTable),
	).From(names.TableDocumentKop).InnerJoin(
		mustache.Render(
			"{{segmenttabel}} ON {{koptabel}}.{{bedrijfsnummer}} = {{segmenttabel}}.{{bedrijfsnummer}} AND {{koptabel}}.{{documentnummer}} = {{segmenttabel}}.{{documentnummer}} AND {{koptabel}}.{{boekjaar}} = {{segmenttabel}}.{{boekjaar}}",
			nameTable,
		),
	).Where(filter.Build()).ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to construct SQL query: %w", err)
	}

	rows, err := db.Query(query, arguments...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	documents := []*model.Document{}
	table := make(map[string]*model.Document)
	for rows.Next() {
		var bedrijfsnummer string
		var documentnummer string
		var boekjaar string
		var documentsoort string
		var documentdatum string
		var boekingsdatum string
		var boekmaand string
		var invoerdatum string
		var invoertijd string
		var boekingsregelnummer string
		var boekingsregelidentificatie string
		var vereffeningsdatum string
		var vereffeningsinvoerdatum string
		var vereffeningsdocument string
		var boekingssleutel string

		if err := rows.Scan(&bedrijfsnummer, &documentnummer, &boekjaar, &documentsoort, &documentdatum, &boekingsdatum, &boekmaand, &invoerdatum, &invoertijd, &boekingsregelnummer, &boekingsregelidentificatie, &vereffeningsdatum, &vereffeningsinvoerdatum, &vereffeningsdocument, &boekingssleutel); err != nil {
			return nil, err
		}

		key := fmt.Sprintf("%s|%s|%s", bedrijfsnummer, documentnummer, boekjaar)

		document, found := table[key]
		if !found {
			parsedDocumentDatum, err := model.ParseYYYYMMSS(documentdatum)
			if err != nil {
				return nil, fmt.Errorf("failed to parse documentdatum: %w", err)
			}

			parsedBoekingsDatum, err := model.ParseYYYYMMSS(boekingsdatum)
			if err != nil {
				return nil, fmt.Errorf("failed to parse boekingsdatum: %w", err)
			}

			parsedInvoerDatum, err := model.ParseYYYYMMSS(invoerdatum)
			if err != nil {
				return nil, fmt.Errorf("failed to parse invoerdatum: %w", err)
			}

			parsedInvoerTijd, err := model.ParseHHMMSS(invoertijd)
			if err != nil {
				return nil, fmt.Errorf("failed to parse invoertijd: %w", err)
			}

			document = &model.Document{
				Bedrijfsnummer: model.NewBedrijfsnummer(bedrijfsnummer),
				DocumentNummer: model.NewDocumentNummer(documentnummer),
				BoekJaar:       model.NewBoekJaar(boekjaar),
				DocumentSoort:  model.NewDocumentSoort(documentsoort),
				DocumentDatum:  parsedDocumentDatum,
				BoekingsDatum:  parsedBoekingsDatum,
				Boekmaand:      model.BoekMaand(boekmaand),
				InvoerDatum:    parsedInvoerDatum,
				InvoerTijd:     parsedInvoerTijd,
				Segmenten:      nil,
			}

			table[key] = document
			documents = append(documents, document)
		}

		parsedBoekingsregelnummer, err := util.ParseInt(boekingsregelnummer)
		if err != nil {
			return nil, fmt.Errorf("failed to parse boekingsregelnummer: %w", err)
		}

		parsedVereffeningsdatum, err := model.ParseYYYYMMSS(vereffeningsdatum)
		if err != nil {
			return nil, fmt.Errorf("failed to parse vereffeningsdatum")
		}

		parsedVereffeningsinvoerdatum, err := model.ParseYYYYMMSS(vereffeningsinvoerdatum)
		if err != nil {
			return nil, fmt.Errorf("failed to parse vereffeningsinvoerdatum")
		}

		segment := model.Segment{
			Regelnummer:               parsedBoekingsregelnummer,
			RegelIdentificatie:        boekingsregelidentificatie,
			VereffeningDatum:          parsedVereffeningsdatum,
			VereffeningInvoerDatum:    parsedVereffeningsinvoerdatum,
			VereffeningDocumentNummer: model.NewDocumentNummer(vereffeningsdocument),
			BoekingSleutel:            model.BoekingSleutel(boekingssleutel),
		}

		document.Segmenten = append(document.Segmenten, segment)
	}

	return documents, nil
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
