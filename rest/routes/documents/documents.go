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

	if bedrijfsnummerString := endpoint.context.Query("bedrijf"); len(bedrijfsnummerString) > 0 {
		bedrijfsNummer, err := model.ParseBedrijfsnummer(bedrijfsnummerString)
		if err != nil {
			return nil, fmt.Errorf("invalid query parameter for bedrijf: %w", err)
		}

		query.WithBedrijfsnummer(bedrijfsNummer)
	}

	return query, nil
}
