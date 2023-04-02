package api

import "net/http"

type RegisterHandlerAPI interface {
	RegisterHandler(writer http.ResponseWriter, requests *http.Request)
	ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request)
}
