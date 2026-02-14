package queries

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"bass-backend/util"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/hoisie/mustache"
)

func ListDocuments() *ListDocumentsQuery {
	return &ListDocumentsQuery{}
}

type ListDocumentsQuery struct {
	whereClauses []squirrel.Sqlizer
}

func (query *ListDocumentsQuery) Execute(db *sql.DB) ([]*model.Document, error) {
	sqlQuery, arguments, err := query.buildSQLQuery()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sqlQuery, arguments...)
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

func (query *ListDocumentsQuery) buildSQLQuery() (string, []any, error) {
	nameTable := map[string]string{
		"koptabel":                   meta.DocumentKop.Table,
		"segmenttabel":               meta.DocumentSegment.Table,
		"bedrijfsnummer":             meta.DocumentKop.BedrijfsNummer,
		"documentnummer":             meta.DocumentKop.DocumentNummer,
		"boekjaar":                   meta.DocumentKop.BoekJaar,
		"documentsoort":              meta.DocumentKop.DocumentSoort,
		"documentdatum":              meta.DocumentKop.DocumentDatum,
		"boekingsdatum":              meta.DocumentKop.BoekingDatum,
		"boekmaand":                  meta.DocumentKop.BoekMaand,
		"invoerdatum":                meta.DocumentKop.InvoerDatum,
		"invoertijd":                 meta.DocumentKop.InvoerTijd,
		"boekingsregelnummer":        meta.DocumentSegment.BoekingsregelNummer,
		"boekingsregelidentificatie": meta.DocumentSegment.BoekingRegelID,
		"vereffeningsdatum":          meta.DocumentSegment.VereffeningDatum,
		"vereffeningsinvoerdatum":    meta.DocumentSegment.VereffeningInvoerDatum,
		"vereffeningsdocument":       meta.DocumentSegment.VereffeningsDocumentNummer,
		"boekingssleutel":            meta.DocumentSegment.Boekingssleutel,
	}

	builder := squirrel.Select(
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
	).From(meta.DocumentKop.Table).InnerJoin(
		mustache.Render(
			"{{segmenttabel}} ON {{koptabel}}.{{bedrijfsnummer}} = {{segmenttabel}}.{{bedrijfsnummer}} AND {{koptabel}}.{{documentnummer}} = {{segmenttabel}}.{{documentnummer}} AND {{koptabel}}.{{boekjaar}} = {{segmenttabel}}.{{boekjaar}}",
			nameTable,
		),
	)

	for _, whereClause := range query.whereClauses {
		builder = builder.Where(whereClause)
	}

	sqlQuery, arguments, err := builder.ToSql()

	return sqlQuery, arguments, err
}

func (query *ListDocumentsQuery) WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer) {
	query.addWhereClause(squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.BedrijfsNummer): bedrijfsnummer.String(),
	})
}

func (query *ListDocumentsQuery) WithBoekjaar(boekjaar model.BoekJaar) {
	query.addWhereClause(squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.BoekJaar): boekjaar.String(),
	})
}

func (query *ListDocumentsQuery) WithDocumentNummerBetween(lower string, upper string) {
	clause := squirrel.Expr(
		fmt.Sprintf("%s.%s BETWEEN ? AND ?", meta.DocumentKop.Table, meta.DocumentKop.DocumentNummer),
		lower,
		upper,
	)

	query.addWhereClause(clause)
}

func (query *ListDocumentsQuery) addWhereClause(clause squirrel.Sqlizer) {
	query.whereClauses = append(query.whereClauses, clause)
}
