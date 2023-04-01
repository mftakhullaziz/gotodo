package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TaskHandlerAPI interface {
	CreateTaskHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
