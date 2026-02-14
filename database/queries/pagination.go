package queries

type Pagination struct {
	Offset   int
	Limit    int
	maxLimit int
}

func InitPagination(maxLimit int) Pagination {
	return Pagination{
		Offset:   0,
		Limit:    maxLimit,
		maxLimit: maxLimit,
	}
}

func (pagination *Pagination) WithLimit(limit int) {
	if limit < 1 {
		limit = 1
	}

	if limit > pagination.maxLimit {
		limit = pagination.maxLimit
	}

	pagination.Limit = limit
}

func (pagination *Pagination) WithOffset(offset int) {
	if offset < 0 {
		offset = 0
	}

	pagination.Offset = offset
}
