package helpers

import (
	"gotodo/internal/domain/models/response"
	"net/http"
)

func CreateResponses(handler interface{}, statusCode int, message string) response.DefaultServiceResponse {
	if HasValue(handler) {
		return response.DefaultServiceResponse{
			StatusCode: statusCode,
			Message:    message,
			IsSuccess:  true,
			Data:       handler,
		}
	}

	var emptyInterface interface{}
	return response.DefaultServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "Request is failed",
		IsSuccess:  false,
		Data:       emptyInterface,
	}
}
