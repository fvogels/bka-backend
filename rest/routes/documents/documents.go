package documents

import (
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
