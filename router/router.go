package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gotodo/apis"
	"gotodo/internal/ports/handlers/api"
	"net/http"
)

var endpoints []string

// NewRouter function to define all router
func NewRouter(login *api.LoginHandlers, register *api.RegisterHandlers, user *api.UserHandlers, tasks *api.TaskHandlers) *httprouter.Router {
	// Init Http Router
	router := httprouter.New()
	// Call url api
	url := apis.Rest()
	// Define all route endpoint and check if nil
	if login != nil {
		LoginRouter(*login, url, router)
	}
	if register != nil {
		RegisterRouter(*register, url, router)
	}
	if tasks != nil {
		TasksRouter(*tasks, url, router)
	}
	if user != nil {
		UserRouter(*user, url, router)
	}
	// Logger url api list
	// utils.ListEndpoints(endpoints)

	return router
}

// LoginRouter function to register endpoint login and logout
func LoginRouter(handler api.LoginHandlers, endpoint apis.Endpoint, router *httprouter.Router) {
	AddedRoute(http.MethodPost, endpoint.AuthenticateLogin, handler.LoginHandler, router)
	AddedRoute(http.MethodPost, endpoint.AuthenticateLogout, handler.LogoutHandler, router)
}

// RegisterRouter function to register endpoint register new user
func RegisterRouter(handler api.RegisterHandlers, endpoint apis.Endpoint, router *httprouter.Router) {
	AddedRoute(http.MethodPost, endpoint.AuthenticateRegister, handler.RegisterHandler, router)
}

// UserRouter function to register endpoint user detail
func UserRouter(handler api.UserHandlers, endpoint apis.Endpoint, router *httprouter.Router) {
	AddedRoute(http.MethodGet, endpoint.AccountUserFind, handler.FindDataUserDetailHandler, router)
	AddedRoute(http.MethodPost, endpoint.AccountUserEdit, handler.UpdateUserDetailHandler, router)
}

// TasksRouter function to register endpoint task
func TasksRouter(handler api.TaskHandlers, endpoint apis.Endpoint, router *httprouter.Router) {
	AddedRoute(http.MethodPost, endpoint.TaskCreate, handler.CreateTaskHandler, router)
	AddedRoute(http.MethodPut, endpoint.TaskUpdate, handler.UpdateTaskHandler, router)
	AddedRoute(http.MethodGet, endpoint.TaskFindByID, handler.FindTaskHandlerById, router)
	AddedRoute(http.MethodGet, endpoint.TaskFind, handler.FindTaskHandler, router)
	AddedRoute(http.MethodPut, endpoint.TaskUpdateStatus, handler.UpdateTaskStatusHandler, router)
	AddedRoute(http.MethodDelete, endpoint.TaskDelete, handler.DeleteTaskHandler, router)
}

// AddedRoute function to register all apis on service http router
func AddedRoute(method, path string, handler func(writer http.ResponseWriter, request *http.Request, params httprouter.Params), router *httprouter.Router) {
	endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
	router.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		handler(writer, request, params)
	})
}
