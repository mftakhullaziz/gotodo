package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	// Initiate Go Mux Router
	router := mux.NewRouter()
	router.Use(helpers.LoggingMiddleware)

	// Initiate prefix url endpoint
	authentications := router.PathPrefix("/api/v1/account").Subrouter()
	authentications.HandleFunc("/register", accountHandler.RegisterHandler).Methods(http.MethodPost)
	authentications.HandleFunc("/login", loginHandler.LoginHandler).Methods(http.MethodPost)
	authentications.HandleFunc("/logout", loginHandler.LogoutHandler).Methods(http.MethodPost)

	users := router.PathPrefix("/api/v1/users").Subrouter()
	users.HandleFunc("/findUser", userDetailHandler.FindDataUserDetailHandler).Methods(http.MethodGet)
	users.HandleFunc("/editUser", userDetailHandler.UpdateUserDetailHandler).Methods(http.MethodPost)

	tasks := router.PathPrefix("/api/v1/task").Subrouter()
	tasks.HandleFunc("/createTask", taskHandler.CreateTaskHandler).Methods(http.MethodPost)
	tasks.HandleFunc("/updateTask/{taskId}", taskHandler.UpdateTaskHandler).Methods(http.MethodPut)
	tasks.HandleFunc("/findTaskId/{taskId}", taskHandler.FindTaskHandlerById).Methods(http.MethodGet)
	tasks.HandleFunc("/findTask", taskHandler.FindTaskHandler).Methods(http.MethodGet)
	tasks.HandleFunc("/deleteTask", taskHandler.DeleteTaskHandler).Methods(http.MethodDelete)
	tasks.HandleFunc("/updateCompletedTask", taskHandler.UpdateTaskStatusHandler).Methods(http.MethodPut)

	// Added logger mux router
	helpers.LogRoutes(router)

	// Server listener
	server := http.Server{Addr: "127.0.0.1:3000", Handler: router}
	err = server.ListenAndServe()
	helpers.PanicIfError(err)
}
