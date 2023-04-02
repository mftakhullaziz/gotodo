package main

import (
	"context"
	"fmt"
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
	log := helpers.LoggerParent()

	ctx := context.Background()
	envName := config.LoadEnv(".env")
	db, errs := database.NewDatabaseConnection(ctx, envName)
	helpers.PanicIfError(errs)
	helpers.LoggerQueryInit(db)

	err := database.MigrateDatabase(db, &record.TaskRecord{}, &record.AccountRecord{}, &record.UserDetailRecord{})
	helpers.PanicIfError(err)

	validate := validator.New()

	// Init Task Handler
	taskRepository := tasksRepository.NewTaskRepositoryImpl(db, validate)
	taskService := tasksService.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := tasksUsecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := tasksHandler.NewTaskHandlerAPI(taskUsecase)

	// Init Account Handler
	userRepository := accountsRepository.NewUserDetailRepositoryImpl(db, validate)
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	accountService := accountsService.NewRegisterServiceImpl(accountRepository, userRepository, validate)
	accountUsecase := accountsUsecase.NewRegisterUseCaseImpl(accountService, validate)
	accountHandler := accountsHandler.NewRegisterHandlerAPI(accountUsecase)

	router := mux.NewRouter()
	router.Use(helpers.LoggingMiddleware)

	handlerTask := router.PathPrefix("/api/v1/task/").Subrouter()
	handlerTask.HandleFunc("/createTask", taskHandler.CreateTaskHandler).Methods("POST")
	handlerTask.HandleFunc("/updateTask/{id}", taskHandler.UpdateTaskHandler).Methods("PUT")
	handlerTask.HandleFunc("/findTaskId/{id}", taskHandler.FindTaskHandlerById).Methods("GET")
	handlerTask.HandleFunc("/findTask", taskHandler.FindTaskHandler).Methods("GET")
	handlerTask.HandleFunc("/deleteTask", taskHandler.DeleteTaskHandler).Methods("DELETE")

	handlerAccount := router.PathPrefix("/api/v1/account/").Subrouter()
	handlerAccount.HandleFunc("/register", accountHandler.RegisterHandler).Methods(http.MethodPost)

	helpers.LogRoutes(router)

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: router,
	}

	log.Infoln("Server Run: ", server.Addr)
	err = server.ListenAndServe()
	helpers.PanicIfError(err)
	fmt.Println("==== SERVER RUNNING ====")
}
