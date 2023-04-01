package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gotodo/config"
	"gotodo/config/database"
	handler "gotodo/internal/adapters/handlers/tasks"
	repository "gotodo/internal/adapters/repositories/tasks"
	service "gotodo/internal/adapters/services/tasks"
	usecase "gotodo/internal/adapters/usecases/tasks"
	"gotodo/internal/helpers"
	"net/http"
)

func main() {
	ctx := context.Background()
	envName := config.LoadEnv(".env")
	db, err := database.NewDatabaseConnection(ctx, envName)
	helpers.PanicIfError(err)

	validate := validator.New()

	taskRepository := repository.NewTaskRepositoryImpl(db, validate)
	taskService := service.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := usecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := handler.NewTaskHandlerAPI(taskUsecase)

	router := httprouter.New()
	router.POST("/handler/api/createTask", taskHandler.CreateTaskHandler)

	router.PanicHandler = helpers.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	fmt.Println(server.Addr)

	err = server.ListenAndServe()
	helpers.PanicIfError(err)
}
