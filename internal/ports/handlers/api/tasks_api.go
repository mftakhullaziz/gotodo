package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TaskHandlerAPI interface {
	CreateTaskHandler(writer http.ResponseWriter, request *http.Request)
	UpdateTaskHandler(writer http.ResponseWriter, request *http.Request, ps httprouter.Params)
	FindTaskHandlerById(writer http.ResponseWriter, request *http.Request, ps httprouter.Params)
	FindTaskHandler(writer http.ResponseWriter, request *http.Request)
	DeleteTaskHandler(writer http.ResponseWriter, request *http.Request)
	UpdateTaskStatusHandler(writer http.ResponseWriter, request *http.Request)
}
