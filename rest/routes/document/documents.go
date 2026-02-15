package document

import (
	"bass-backend/database/queries"
	"bass-backend/model"
	"bass-backend/util"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	MaximumDocumentsPerRequest = 100
)

type documentEndpoint struct {
	context  *gin.Context
	database *sql.DB
}

type CountResponse struct {
	Count int `json:"count"`
}

type ListResponse struct {
	Documents []*model.Document `json:"documents"`
}

type paginationInformation struct {
	Limit  int
	Offset int
}

type query interface {
	WithBoekjaar(model.BoekJaar)
	WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer)
	WithDocumentNummerBetween(lower string, upper string)
}

func Handle(database *sql.DB, context *gin.Context) {
	endpoint := documentEndpoint{
		context:  context,
		database: database,
	}

	endpoint.execute()
}

func (endpoint *documentEndpoint) execute() {
	context := endpoint.context

	paginationInformation, err := endpoint.extractLimitAndOffsetFromQueryParameters()
	if err != nil {
		context.String(http.StatusBadRequest, "invalid limit/offset query parameters")
		return
	}

	if paginationInformation == nil {
		endpoint.respondWithCount()
		return
	} else {
		endpoint.respondWithList(paginationInformation)
		return
	}
}

func (endpoint *documentEndpoint) respondWithCount() {
	context := endpoint.context

	query, err := endpoint.buildCountQuery()
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	count, err := query.Execute(endpoint.database)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	response := CountResponse{
		Count: count,
	}

	context.JSON(http.StatusOK, response)
}

func (endpoint *documentEndpoint) respondWithList(paginationInformation *paginationInformation) {
	context := endpoint.context

	query, err := endpoint.buildListQuery(paginationInformation)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	documents, err := query.Execute(endpoint.database)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	response := ListResponse{
		Documents: documents,
	}

	context.JSON(http.StatusOK, response)
}

func (endpoint *documentEndpoint) extractLimitAndOffsetFromQueryParameters() (*paginationInformation, error) {
	context := endpoint.context

	limitString := context.Query("limit")
	offsetString := context.Query("offset")

	if len(limitString) == 0 && len(offsetString) == 0 {
		return nil, nil
	}

	if len(limitString) == 0 {
		limitString = strconv.FormatInt(MaximumDocumentsPerRequest, 10)
	}

	if len(offsetString) == 0 {
		offsetString = "0"
	}

	limit, err := util.ParseInt(limitString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limit %s: %w", limit, err)
	}

	offset, err := util.ParseInt(offsetString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse offset %s: %w", offset, err)
	}

	return &paginationInformation{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (endpoint *documentEndpoint) buildCountQuery() (*queries.CountDocumentsQuery, error) {
	query := queries.CountDocuments()

	endpoint.processBedrijfQueryParameter(query)
	endpoint.processBoekjaarQueryParameter(query)
	endpoint.processDocumentnummerIntervalQueryParameter(query)

	return query, nil
}

func (endpoint *documentEndpoint) buildListQuery(pagination *paginationInformation) (*queries.ListDocumentsQuery, error) {
	query := queries.ListDocuments()

	endpoint.processBedrijfQueryParameter(query)
	endpoint.processBoekjaarQueryParameter(query)
	endpoint.processDocumentnummerIntervalQueryParameter(query)
	endpoint.processPagination(query, pagination)

	return query, nil
}

func (endpoint *documentEndpoint) processBedrijfQueryParameter(query query) error {
	if bedrijfsnummerString := endpoint.context.Query("bedrijfsnummer"); len(bedrijfsnummerString) > 0 {
		bedrijfsNummer, err := model.ParseBedrijfsnummer(bedrijfsnummerString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijfsnummer: %w", err)
		}

		query.WithBedrijfsnummer(bedrijfsNummer)
	}

	return nil
}

func (endpoint *documentEndpoint) processBoekjaarQueryParameter(query query) error {
	if boekjaarString := endpoint.context.Query("boekjaar"); len(boekjaarString) > 0 {
		boekjaar, err := model.ParseBoekJaar(boekjaarString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBoekjaar(boekjaar)
	}

	return nil
}

func (endpoint *documentEndpoint) processDocumentnummerIntervalQueryParameter(query query) error {
	if documentnummerIntervalString := endpoint.context.Query("document"); len(documentnummerIntervalString) > 0 {
		bounds := strings.Split(documentnummerIntervalString, "-")

		if len(bounds) != 2 {
			return fmt.Errorf("invalid documentnummer interval: %s", documentnummerIntervalString)
		}

		query.WithDocumentNummerBetween(bounds[0], bounds[1])
	}

	return nil
}

func (endpoint *documentEndpoint) processPagination(query *queries.ListDocumentsQuery, pagination *paginationInformation) {
	query.WithLimitAndOffset(pagination.Limit, pagination.Offset)
}
