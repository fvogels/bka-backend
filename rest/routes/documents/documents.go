package documents

import (
	"bass-backend/database/queries"
	"bass-backend/model"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type listDocumentEndpoint struct {
	context  *gin.Context
	database *sql.DB
}

type Response struct {
	Count int `json:"count"`
}

type query interface {
	WithBoekjaar(model.BoekJaar)
	WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer)
	WithDocumentNummerBetween(lower string, upper string)
}

func Handle(database *sql.DB, context *gin.Context) {
	endpoint := listDocumentEndpoint{
		context:  context,
		database: database,
	}

	endpoint.execute()
}

func (endpoint *listDocumentEndpoint) execute() {
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

	response := Response{
		Count: count,
	}

	context.JSON(http.StatusOK, response)
}

func (endpoint *listDocumentEndpoint) buildCountQuery() (*queries.CountDocumentsQuery, error) {
	query := queries.CountDocuments()

	endpoint.processBedrijfQueryParameter(query)
	endpoint.processBoekjaarQueryParameter(query)
	endpoint.processDocumentnummerIntervalQueryParameter(query)

	return query, nil
}

func (endpoint *listDocumentEndpoint) processBedrijfQueryParameter(query query) error {
	if bedrijfsnummerString := endpoint.context.Query("bedrijf"); len(bedrijfsnummerString) > 0 {
		bedrijfsNummer, err := model.ParseBedrijfsnummer(bedrijfsnummerString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBedrijfsnummer(bedrijfsNummer)
	}

	return nil
}

func (endpoint *listDocumentEndpoint) processBoekjaarQueryParameter(query query) error {
	if boekjaarString := endpoint.context.Query("boekjaar"); len(boekjaarString) > 0 {
		boekjaar, err := model.ParseBoekJaar(boekjaarString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBoekjaar(boekjaar)
	}

	return nil
}

func (endpoint *listDocumentEndpoint) processDocumentnummerIntervalQueryParameter(query query) error {
	if documentnummerIntervalString := endpoint.context.Query("document"); len(documentnummerIntervalString) > 0 {
		bounds := strings.Split(documentnummerIntervalString, "-")

		if len(bounds) != 2 {
			return fmt.Errorf("invalid documentnummer interval: %s", documentnummerIntervalString)
		}

		query.WithDocumentNummerBetween(bounds[0], bounds[1])
	}

	return nil
}
