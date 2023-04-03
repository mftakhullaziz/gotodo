package api

import "net/http"

type LoginHandlerAPI interface {
	LoginHandler(writer http.ResponseWriter, requests *http.Request)
	LogoutHandler(writer http.ResponseWriter, requests *http.Request)
}
