package api

import (
	"net/http"
)

type TaskHandlerAPI interface {
	CreateTaskHandler(writer http.ResponseWriter, request *http.Request)
	UpdateTaskHandler(writer http.ResponseWriter, request *http.Request)
	FindTaskHandlerById(writer http.ResponseWriter, request *http.Request)
	FindTaskHandler(writer http.ResponseWriter, request *http.Request)
	DeleteTaskHandler(writer http.ResponseWriter, request *http.Request)
}
