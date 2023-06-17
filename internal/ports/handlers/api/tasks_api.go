package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TaskHandlers interface {
	CreateTaskHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	UpdateTaskHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindTaskHandlerById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindTaskHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	DeleteTaskHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	UpdateTaskStatusHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
}
