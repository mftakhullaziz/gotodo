package login

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gotodo/config"
	"gotodo/config/database"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	accountsService "gotodo/internal/adapters/services/accounts"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	"gotodo/internal/utils"
	"gotodo/router"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Initiate Setup Test Database
func InitTestDB() *gorm.DB {
	// Initiate context
	ctx := context.Background()
	// The param path is dynamic following testing path file
	env := config.LoadEnvFromFile("../../../../../..")
	// Do function database connection
	db, _ := database.NewDatabaseConnection(ctx, env)

	return db
}

func InitRouter(db *gorm.DB) *httprouter.Router {
	// Initiate validator
	validate := validator.New()
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := NewLoginHandlers(loginUsecase)

	r := router.NewRouter(&loginHandler, nil, nil, nil)
	return r
}

func TestLoginHandlerAPI_LoginHandler(t *testing.T) {
	db := InitTestDB()
	initRouter := InitRouter(db)

	t.Run("Login Test Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username" : "johndoe_416", "password": "password"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		initRouter.ServeHTTP(recorder, httpRequest)

		data := recorder.Result()
		body, _ := io.ReadAll(data.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		// Verify response
		statusCode, _ := utils.ValueToInt(responseBody["status_code"])
		assert.Equalf(t, statusCode, http.StatusCreated, "Expected: %d, but got: %d", statusCode, http.StatusCreated)
		assert.Equalf(t, responseBody["message"], "login account successfully!", "Expected: %s, but got: %s", responseBody["message"], "login account successfully!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})

	t.Run("Login Test Is Failed", func(t *testing.T) {
		// Scenario failed is username or password is not correct
		requestBody := strings.NewReader(`{"username" : "johndoe_416", "password": "failed"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		initRouter.ServeHTTP(recorder, httpRequest)

		data := recorder.Result()
		body, _ := io.ReadAll(data.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		// Verify response
		statusCode, _ := utils.ValueToInt(responseBody["status_code"])
		assert.Equalf(t, statusCode, http.StatusBadRequest, "Expected: %d, but got: %d", statusCode, http.StatusBadRequest)
		assert.Equalf(t, responseBody["message"], "username and password not valid please check again!", "Expected: %s, but got: %s", responseBody["message"], "username and password not valid please check again!")
		assert.Equalf(t, responseBody["is_success"], false, "Expected: %s, but got: %s", responseBody["is_success"], false)

		// Verify response data
		assert.Emptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.Nilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}
