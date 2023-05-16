package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gotodo/config"
	"gotodo/config/database"
	accountsHandler "gotodo/internal/adapters/handlers/accounts"
	tasksHandler "gotodo/internal/adapters/handlers/tasks"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	tasksRepository "gotodo/internal/adapters/repositories/tasks"
	accountsService "gotodo/internal/adapters/services/accounts"
	tasksService "gotodo/internal/adapters/services/tasks"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	tasksUsecase "gotodo/internal/adapters/usecases/tasks"
	"gotodo/internal/helpers"
	"gotodo/internal/persistence/record"
	"net/http"
)

type EndpointPrefix struct {
	authenticate string
	account      string
	task         string
}

type Endpoint struct {
	authenticateRegister string
	authenticateLogin    string
	authenticateLogout   string
	accountUserFind      string
	accountUserEdit      string
	taskCreate           string
	taskUpdate           string
	taskUpdateStatus     string
	taskFindByID         string
	taskFind             string
	taskDelete           string
}

var endpoints []string

func main() {
	// Initiate context
	ctx := context.Background()
	envName := config.LoadEnv(".env")

	// Do function database connection
	db, errs := database.NewDatabaseConnection(ctx, envName)
	helpers.PanicIfErrorWithCustomMessage(errs, "new database connection is failed")
	helpers.LoggerQueryInit(db)

	// Do migration database
	err := database.MigrateDatabase(db,
		&record.TaskRecord{},
		&record.AccountRecord{},
		&record.UserDetailRecord{},
		&record.AccountLoginHistoriesRecord{})
	helpers.LoggerIfError(err)

	// Initiate validator
	validate := validator.New()

	// Init Task Handler
	taskRepository := tasksRepository.NewTaskRepositoryImpl(db, validate)
	taskService := tasksService.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := tasksUsecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := tasksHandler.NewTaskHandlerAPI(taskUsecase)

	// Init Register Handler
	userRepository := accountsRepository.NewUserDetailRepositoryImpl(db, validate)
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	accountService := accountsService.NewRegisterServiceImpl(accountRepository, userRepository, validate)
	accountUsecase := accountsUsecase.NewRegisterUseCaseImpl(accountService, validate)
	accountHandler := accountsHandler.NewRegisterHandlerAPI(accountUsecase)

	// Login Handler
	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := accountsHandler.NewLoginHandlerAPI(loginUsecase)

	// User Detail Handler
	userDetailService := accountsService.NewUserDetailServiceImpl(userRepository, accountRepository, validate)
	userDetailUsecase := accountsUsecase.NewUserDetailUsecaseImpl(userDetailService, validate)
	userDetailHandler := accountsHandler.NewUserDetailHandlerAPI(userDetailUsecase)

	router := httprouter.New()

	apiRoute := EndpointPrefix{
		authenticate: "/api/v1/authenticate/",
		account:      "/api/v1/user/",
		task:         "/api/v1/task/",
	}

	api := Endpoint{
		authenticateRegister: apiRoute.authenticate + "register",
		authenticateLogin:    apiRoute.authenticate + "login",
		authenticateLogout:   apiRoute.authenticate + "logout",
		accountUserFind:      apiRoute.account + "find",
		accountUserEdit:      apiRoute.account + "edit",
		taskCreate:           apiRoute.task + "create",
		taskUpdate:           apiRoute.task + "update/:task_id",
		taskFindByID:         apiRoute.task + "find/:task_id",
		taskFind:             apiRoute.task + "find",
		taskDelete:           apiRoute.task + "delete",
		taskUpdateStatus:     apiRoute.task + "update_status",
	}

	RegisterEndpoint(http.MethodPost, api.authenticateRegister, accountHandler.RegisterHandler, router)
	RegisterEndpoint(http.MethodPost, api.authenticateLogin, loginHandler.LoginHandler, router)
	RegisterEndpoint(http.MethodPost, api.authenticateLogout, loginHandler.LogoutHandler, router)
	RegisterEndpoint(http.MethodGet, api.accountUserFind, userDetailHandler.FindDataUserDetailHandler, router)
	RegisterEndpoint(http.MethodPost, api.accountUserEdit, userDetailHandler.UpdateUserDetailHandler, router)

	RegisterEndpoint(http.MethodPost, api.taskCreate, taskHandler.CreateTaskHandler, router)
	RegisterEndpoint(http.MethodPut, api.taskUpdate, taskHandler.UpdateTaskHandler, router)
	RegisterEndpoint(http.MethodGet, api.taskFindByID, taskHandler.FindTaskHandlerById, router)
	RegisterEndpoint(http.MethodGet, api.taskFind, taskHandler.FindTaskHandler, router)
	RegisterEndpoint(http.MethodDelete, api.taskDelete, taskHandler.DeleteTaskHandler, router)
	RegisterEndpoint(http.MethodPut, api.taskUpdateStatus, taskHandler.UpdateTaskStatusHandler, router)

	// Logger endpoint list
	EndpointList()
	loggedRouter := loggingMiddleware(router)

	// Server listener
	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: loggedRouter,
	}

	ok := server.ListenAndServe()
	helpers.PanicIfError(ok)
}

// RegisterEndpoint function to register all endpoint on service http router
func RegisterEndpoint(method, path string, handler func(writer http.ResponseWriter, requests *http.Request), router *httprouter.Router) {
	endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
	router.HandlerFunc(method, path, handler)
}

func EndpointList() {
	loggerParent := helpers.LoggerParent()
	loggerParent.Infoln("Registered endpoints:")
	for _, endpoint := range endpoints {
		loggerParent.Infoln(endpoint)
	}
}

// Middleware function to log requests and responses
func loggingMiddleware(next http.Handler) http.Handler {
	loggerParent := helpers.LoggerParent()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		loggerParent.Infof("Received request: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log response details
		loggerParent.Infof("Sent response: %s %s", r.Method, r.URL.Path)
	})
}
