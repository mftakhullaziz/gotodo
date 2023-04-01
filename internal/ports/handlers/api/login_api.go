package api

import "net/http"

type LoginHandlerAPI interface {
	LoginHandler(writer http.ResponseWriter, request *http.Request)
	LogoutHandler(writer http.ResponseWriter, request *http.Request)
}
