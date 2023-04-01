package api

import "net/http"

type RegisterHandlerAPI interface {
	RegisterHandler(writer http.ResponseWriter, request *http.Request)
	ForgotPasswordHandler(writer http.ResponseWriter, request *http.Request)
}
