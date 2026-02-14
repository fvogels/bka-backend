package documents

import (
	"bass-backend/database/queries"
	"bass-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listDocumentEndpoint struct {
	context *gin.Context
}

type Response struct {
	Count int `json:"count"`
}

type query interface {
	WithBoekjaar(model.BoekJaar)
	WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer)
	WithDocumentNummerBetween(lower string, upper string)
}

func Handle(context *gin.Context) {
	endpoint := listDocumentEndpoint{
		context: context,
	}

	endpoint.execute()
}

func (endpoint *listDocumentEndpoint) execute() {
	response := Response{
		Count: 10,
	}

	endpoint.context.JSON(http.StatusOK, response)
}

func (endpoint *listDocumentEndpoint) buildCountQuery() (*queries.CountDocumentsQuery, error) {
	query := queries.CountDocuments()

	endpoint.parseBedrijfQueryParameter(query)
	endpoint.parseBoekjaarQueryParameter(query)

	return query, nil
}

func (endpoint *listDocumentEndpoint) parseBedrijfQueryParameter(query query) error {
	if bedrijfsnummerString := endpoint.context.Query("bedrijf"); len(bedrijfsnummerString) > 0 {
		bedrijfsNummer, err := model.ParseBedrijfsnummer(bedrijfsnummerString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBedrijfsnummer(bedrijfsNummer)
	}

	return nil
}

func (endpoint *listDocumentEndpoint) parseBoekjaarQueryParameter(query query) error {
	if boekjaarString := endpoint.context.Query("boekjaar"); len(boekjaarString) > 0 {
		boekjaar, err := model.ParseBoekJaar(boekjaarString)
		if err != nil {
			return fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBoekjaar(boekjaar)
	}

	return nil
}
