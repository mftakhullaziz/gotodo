package api

import "net/http"

type UserHandlerAPI interface {
	UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request)
	FindDataUserDetailHandler(writer http.ResponseWriter, requests *http.Request)
	DeleteUserHandler(writer http.ResponseWriter, requests *http.Request)
	FindAllUserAccountHandler(writer http.ResponseWriter, requests *http.Request)
}
