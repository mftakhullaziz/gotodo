package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gotodo/api"
	handler "gotodo/internal/ports/handlers/api"
	"gotodo/internal/utils"
	"net/http"
)

var endpoints []string

// NewRouter function to define all router
func NewRouter(
	login *handler.LoginHandlers,
	register *handler.RegisterHandlers,
	user *handler.UserHandlers,
	tasks *handler.TaskHandlers) *httprouter.Router {

	// Init Http Router
	router := httprouter.New()

	// Define all route endpoint and check if nil
	if login != nil {
		LoginRouter(*login, router)
	}
	if register != nil {
		RegisterRouter(*register, router)
	}
	if tasks != nil {
		TasksRouter(*tasks, router)
	}
	if user != nil {
		UserRouter(*user, router)
	}
	// Logger url api list
	utils.ListEndpoints(endpoints)

	return router
}

// LoginRouter function to register endpoint login and logout
func LoginRouter(handler handler.LoginHandlers, router *httprouter.Router) {
	AddedRoute(http.MethodPost, api.Login, handler.LoginHandler, router)
	AddedRoute(http.MethodPost, api.Logout, handler.LogoutHandler, router)
}

// RegisterRouter function to register endpoint register new user
func RegisterRouter(handler handler.RegisterHandlers, router *httprouter.Router) {
	AddedRoute(http.MethodPost, api.Register, handler.RegisterHandler, router)
}

// UserRouter function to register endpoint user detail
func UserRouter(handler handler.UserHandlers, router *httprouter.Router) {
	AddedRoute(http.MethodGet, api.UserFind, handler.FindDataUserDetailHandler, router)
	AddedRoute(http.MethodPost, api.UserEdit, handler.UpdateUserDetailHandler, router)
}

// TasksRouter function to register endpoint task
func TasksRouter(handler handler.TaskHandlers, router *httprouter.Router) {
	AddedRoute(http.MethodPost, api.TaskCreate, handler.CreateTaskHandler, router)
	AddedRoute(http.MethodPut, api.TaskUpdate, handler.UpdateTaskHandler, router)
	AddedRoute(http.MethodGet, api.TaskFindById, handler.FindTaskHandlerById, router)
	AddedRoute(http.MethodGet, api.TaskFind, handler.FindTaskHandler, router)
	AddedRoute(http.MethodPut, api.TaskUpdateStatus, handler.UpdateTaskStatusHandler, router)
	AddedRoute(http.MethodDelete, api.TaskDelete, handler.DeleteTaskHandler, router)
}

// AddedRoute function to register all api on service http router
func AddedRoute(method, path string, handler func(writer http.ResponseWriter, request *http.Request, params httprouter.Params), router *httprouter.Router) {
	endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
	router.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		handler(writer, request, params)
	})
}
