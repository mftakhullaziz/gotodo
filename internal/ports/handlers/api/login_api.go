package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LoginHandlerAPI interface {
	LoginHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
	LogoutHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
}
