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
			Data:       handler}
	}

	var emptyInterface interface{}
	return response.DefaultServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    errMessage,
		IsSuccess:  false,
		Data:       emptyInterface}
}

func BuildResponseWithAuthorization(
	handler interface{},
	statusCode int,
	taskId int,
	userId string,
	message1 string) response.DefaultServiceResponse {
	if HasValue(handler) && userId != "" && taskId != 0 {
		return response.DefaultServiceResponse{
			StatusCode: statusCode,
			Message:    message1,
			IsSuccess:  true,
			Data:       handler,
		}
	}

	var emptyInterface interface{}
	return response.DefaultServiceResponse{
		StatusCode: http.StatusOK,
		Message:    message1,
		IsSuccess:  true,
		Data:       emptyInterface,
	}
}

func BuildEmptyResponse(messages string) response.DefaultServiceResponse {
	var emptyInterface interface{}
	return response.DefaultServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    messages,
		IsSuccess:  false,
		Data:       emptyInterface,
	}
}

func BuildAllResponseWithAuthorization(handler interface{},
	message string, totalData int, requestAt string) response.DefaultServiceAllResponse {
	if HasValueSlice(handler) {
		return response.DefaultServiceAllResponse{
			StatusCode: http.StatusOK,
			Message:    message,
			IsSuccess:  true,
			Data:       handler,
			TotalData:  totalData,
			RequestAt:  requestAt,
		}
	}

	var emptyInterface []interface{}
	return response.DefaultServiceAllResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		IsSuccess:  true,
		Data:       emptyInterface,
		TotalData:  0,
		RequestAt:  requestAt,
	}
}
