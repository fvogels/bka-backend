package queries

import (
	"bass-backend/database/meta"
	"bass-backend/database/queries/filters"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func CountDocuments() *CountDocumentsQuery {
	whereClauses := []squirrel.Sqlizer{}

	return &CountDocumentsQuery{
		Bedrijfsnummer:         filters.InitBedrijfsnummerFilter(&whereClauses),
		Boekjaar:               filters.InitBoekjaarFilter(&whereClauses),
		DocumentnummerInterval: filters.InitDocumentnummerInterval(&whereClauses),
	}
}

type CountDocumentsQuery struct {
	filters.Bedrijfsnummer
	filters.Boekjaar
	filters.DocumentnummerInterval

	whereClauses *[]squirrel.Sqlizer
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

	for _, whereClause := range *query.whereClauses {
		builder = builder.Where(whereClause)
	}

	sqlQuery, arguments, err := builder.ToSql()

	return sqlQuery, arguments, err
}
