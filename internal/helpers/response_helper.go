package helpers

import (
	"gotodo/internal/domain/models/response"
	"net/http"
)

func CreateResponses(handler interface{}, statusCode int, message string, errMessage string) response.DefaultServiceResponse {
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
		Message:    errMessage,
		IsSuccess:  false,
		Data:       emptyInterface,
	}
}

func BuildResponseWithAuthorization(handler interface{}, statusCode int, userId string, message string, errMessage string) response.DefaultServiceResponse {
	if HasValue(handler) && userId != "" {
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
		Message:    errMessage,
		IsSuccess:  false,
		Data:       emptyInterface,
	}
}
