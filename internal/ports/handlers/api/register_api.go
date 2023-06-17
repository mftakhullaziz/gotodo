package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type RegisterHandlers interface {
	RegisterHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
	ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params)
}
