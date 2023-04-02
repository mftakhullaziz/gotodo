package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gotodo/config"
	"gotodo/config/database"
	handler "gotodo/internal/adapters/handlers/tasks"
	repository "gotodo/internal/adapters/repositories/tasks"
	service "gotodo/internal/adapters/services/tasks"
	usecase "gotodo/internal/adapters/usecases/tasks"
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

	taskRepository := repository.NewTaskRepositoryImpl(db, validate)
	taskService := service.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := usecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := handler.NewTaskHandlerAPI(taskUsecase)

	router := mux.NewRouter()
	router.Use(helpers.LoggingMiddleware)

	handlerTask := router.PathPrefix("/api/v1/task/").Subrouter()
	handlerTask.HandleFunc("/createTask", taskHandler.CreateTaskHandler).Methods("POST")
	handlerTask.HandleFunc("/updateTask/{id}", taskHandler.UpdateTaskHandler).Methods("PUT")
	handlerTask.HandleFunc("/findTaskId/{id}", taskHandler.FindTaskHandlerById).Methods("GET")
	handlerTask.HandleFunc("/findTask", taskHandler.FindTaskHandler).Methods("GET")
	handlerTask.HandleFunc("/deleteTask", taskHandler.DeleteTaskHandler).Methods("DELETE")

	// handlerAccount := router.PathPrefix("/api/v1/account/").Subrouter()

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
