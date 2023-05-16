package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gotodo/config"
	"gotodo/config/database"
	"gotodo/endpoint"
	accountsHandler "gotodo/internal/adapters/handlers/accounts"
	tasksHandler "gotodo/internal/adapters/handlers/tasks"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	tasksRepository "gotodo/internal/adapters/repositories/tasks"
	accountsService "gotodo/internal/adapters/services/accounts"
	tasksService "gotodo/internal/adapters/services/tasks"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	tasksUsecase "gotodo/internal/adapters/usecases/tasks"
	"gotodo/internal/persistence/record"
	"gotodo/internal/utils"
	"net/http"
)

var endpoints []string

func main() {
	// Initiate context
	ctx := context.Background()
	envName := config.LoadEnv(".env")

	// Do function database connection
	db, err := database.NewDatabaseConnection(ctx, envName)
	utils.PanicIfErrorWithCustomMessage(err, "new database connection is failed")
	utils.LoggerQueryInit(db)

	// Do migration database
	err = database.MigrateDatabase(
		db,
		&record.TaskRecord{},
		&record.AccountRecord{},
		&record.UserDetailRecord{},
		&record.AccountLoginHistoriesRecord{},
	)
	utils.LoggerIfError(err)

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

	// Init Http Router
	router := httprouter.New()
	// Call endpoint
	api := endpoint.Rest()

	// Register all endpoint
	RegisterEndpoint(http.MethodPost, api.AuthenticateRegister, accountHandler.RegisterHandler, router)
	RegisterEndpoint(http.MethodPost, api.AuthenticateLogin, loginHandler.LoginHandler, router)
	RegisterEndpoint(http.MethodPost, api.AuthenticateLogout, loginHandler.LogoutHandler, router)
	RegisterEndpoint(http.MethodGet, api.AccountUserFind, userDetailHandler.FindDataUserDetailHandler, router)
	RegisterEndpoint(http.MethodPost, api.AccountUserEdit, userDetailHandler.UpdateUserDetailHandler, router)
	RegisterEndpoint(http.MethodPost, api.TaskCreate, taskHandler.CreateTaskHandler, router)
	RegisterEndpoint(http.MethodPut, api.TaskUpdate, taskHandler.UpdateTaskHandler, router)
	RegisterEndpoint(http.MethodGet, api.TaskFindByID, taskHandler.FindTaskHandlerById, router)
	RegisterEndpoint(http.MethodGet, api.TaskFind, taskHandler.FindTaskHandler, router)
	RegisterEndpoint(http.MethodDelete, api.TaskDelete, taskHandler.DeleteTaskHandler, router)
	RegisterEndpoint(http.MethodPut, api.TaskUpdateStatus, taskHandler.UpdateTaskStatusHandler, router)

	// Logger endpoint list
	ListEndpoints()
	// Init Router in Logger
	LoggerRouter := utils.LoggerMiddleware(router)

	// Server listener
	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: LoggerRouter,
	}

	ok := server.ListenAndServe()
	utils.PanicIfError(ok)
}

// RegisterEndpoint function to register all endpoint on service http router
func RegisterEndpoint(method, path string,
	handler func(writer http.ResponseWriter, requests *http.Request),
	router *httprouter.Router) {

	endpoints = append(endpoints, fmt.Sprintf("%s %s", method, path))
	router.HandlerFunc(method, path, handler)
}

// ListEndpoints function to show all registered endpoint
func ListEndpoints() {
	log := utils.LoggerParent()
	log.Infoln("Registered endpoints:")
	for _, e := range endpoints {
		log.Infoln(e)
	}
}
