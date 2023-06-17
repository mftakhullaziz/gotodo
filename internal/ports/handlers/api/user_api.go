package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserHandlerAPI interface {
	FindDataUserDetailHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
	UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
	DeleteUserHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
}
