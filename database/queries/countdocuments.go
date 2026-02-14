package queries

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func CountDocuments() *CountDocumentsQuery {
	return &CountDocumentsQuery{}
}

type CountDocumentsQuery struct {
	whereClauses []squirrel.Sqlizer
}

func (query *CountDocumentsQuery) Execute(db *sql.DB) (int, error) {
	sqlQuery, arguments, err := query.buildSQLQuery()
	if err != nil {
		return 0, err
	}

	row := db.QueryRow(sqlQuery, arguments...)

	var documentCount int
	if err := row.Scan(&documentCount); err != nil {
		return 0, fmt.Errorf("failed to count the number of documents: %w", err)
	}

	return documentCount, nil
}

func (query *CountDocumentsQuery) buildSQLQuery() (string, []any, error) {
	builder := squirrel.Select("COUNT(*)").From(meta.DocumentKop.Table)

	for _, whereClause := range query.whereClauses {
		builder = builder.Where(whereClause)
	}

	sqlQuery, arguments, err := builder.ToSql()

	return sqlQuery, arguments, err
}

func (query *CountDocumentsQuery) WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer) {
	query.addWhereClause(squirrel.Eq{
		meta.DocumentKop.BedrijfsNummer: bedrijfsnummer.String(),
	})
}

func (query *CountDocumentsQuery) WithBoekjaar(boekjaar model.BoekJaar) {
	query.addWhereClause(squirrel.Eq{
		meta.DocumentKop.BoekJaar: boekjaar.String(),
	})
}

func (query *CountDocumentsQuery) WithDocumentNummerBetween(lower string, upper string) {
	clause := squirrel.Expr(
		fmt.Sprintf("%s BETWEEN ? AND ?", meta.DocumentKop.DocumentNummer),
		lower,
		upper,
	)

	query.addWhereClause(clause)
}

func (query *CountDocumentsQuery) addWhereClause(clause squirrel.Sqlizer) {
	query.whereClauses = append(query.whereClauses, clause)
}
