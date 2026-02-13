package queries

import (
	"bass-backend/database/filters"
	"bass-backend/database/names"
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
