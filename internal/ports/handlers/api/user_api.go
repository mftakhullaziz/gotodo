package api

import "net/http"

type UserHandlerAPI interface {
	FindDataUserDetailHandler(writer http.ResponseWriter, requests *http.Request)
	UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request)
	DeleteUserHandler(writer http.ResponseWriter, requests *http.Request)
}
