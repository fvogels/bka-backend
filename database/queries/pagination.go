package queries

import "github.com/Masterminds/squirrel"

type Pagination struct {
	Offset *int
	Limit  *int
}

func InitPagination() Pagination {
	return Pagination{
		Offset: nil,
		Limit:  nil,
	}
}

func (pagination *Pagination) WithLimit(limit int) {
	if limit < 1 {
		limit = 1
	}

	pagination.Limit = &limit
}

// withOffset sets the offset.
// It is kept hidden because having an offset without limit is disallowed.
func (pagination *Pagination) withOffset(offset int) {
	if offset < 0 {
		offset = 0
	}

	pagination.Offset = &offset
}

func (pagination *Pagination) WithLimitAndOffset(limit int, offset int) {
	pagination.WithLimit(limit)
	pagination.withOffset(offset)
}

func (pagination *Pagination) Apply(query squirrel.SelectBuilder) squirrel.SelectBuilder {
	if pagination.Offset != nil {
		query = query.Offset(uint64(*pagination.Offset))
	}

	if pagination.Limit != nil {
		query = query.Limit(uint64(*pagination.Limit))
	}

	return query
}
