package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"gotodo/config"
	"gotodo/config/database"
	accountsHandler "gotodo/internal/adapters/handlers/accounts"
	loginsHandler "gotodo/internal/adapters/handlers/accounts/login"
	tasksHandler "gotodo/internal/adapters/handlers/tasks"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	tasksRepository "gotodo/internal/adapters/repositories/tasks"
	accountsService "gotodo/internal/adapters/services/accounts"
	tasksService "gotodo/internal/adapters/services/tasks"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	tasksUsecase "gotodo/internal/adapters/usecases/tasks"
	"gotodo/internal/persistence/record"
	"gotodo/internal/utils"
	"gotodo/router"
	"net/http"
)

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

	// Task Handler
	taskRepository := tasksRepository.NewTaskRepositoryImpl(db, validate)
	taskService := tasksService.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := tasksUsecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := tasksHandler.NewTaskHandlerAPI(taskUsecase)

	// Register Handler
	userRepository := accountsRepository.NewUserDetailRepositoryImpl(db, validate)
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	accountService := accountsService.NewRegisterServiceImpl(accountRepository, userRepository, validate)
	accountUsecase := accountsUsecase.NewRegisterUseCaseImpl(accountService, validate)
	accountHandler := accountsHandler.NewRegisterHandlerAPI(accountUsecase)

	// Login Handler
	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := loginsHandler.NewLoginHandlerAPI(loginUsecase)

	// User Detail Handler
	userDetailService := accountsService.NewUserDetailServiceImpl(userRepository, accountRepository, validate)
	userDetailUsecase := accountsUsecase.NewUserDetailUsecaseImpl(userDetailService, validate)
	userDetailHandler := accountsHandler.NewUserDetailHandlerAPI(userDetailUsecase)

	// Call http router
	app := router.NewRouter(loginHandler, accountHandler, userDetailHandler, taskHandler)

	// Init Router in Logger
	LoggerRouter := utils.LoggerMiddleware(app)

	// Server listener
	ok := http.ListenAndServe(":3000", LoggerRouter)
	utils.PanicIfError(ok)
}
