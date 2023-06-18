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

func TestLoginHandlers_LoginHandler(t *testing.T) {
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
		username := responseBody["data"].(map[string]interface{})["username"].(string)
		id := responseBody["data"].(map[string]interface{})["account_id"].(float64)
		assert.Equalf(t, username, "johndoe_416", "Expected: %s, but got: %s", username, "johndoe_416")
		assert.Equalf(t, int(id), 12, "Expected: %d, but got: %d", int(id), 12)
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

func TestLoginHandlers_LogoutHandler(t *testing.T) {
	db := InitTestDB()
	initRouter := InitRouter(db)

	t.Run("Logout Test From User Login Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username" : "johndoe_416", "password": "password"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		initRouter.ServeHTTP(recorder, httpRequest)

		data := recorder.Result()
		body, _ := io.ReadAll(data.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		// Key from login
		authorizationKey := responseBody["data"].(map[string]interface{})["token"].(string)
		httpRequest = httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/logout", nil)

		// Setup Header key for parsing token in authorization
		headers := make(http.Header)
		headers.Set("Content-Type", "application/json")
		headers.Set("Authorization", authorizationKey)
		httpRequest.Header = headers

		// Execute the request
		recorder = httptest.NewRecorder()
		initRouter.ServeHTTP(recorder, httpRequest)

		data1 := recorder.Result()
		body1, _ := io.ReadAll(data1.Body)
		var responseBody1 map[string]interface{}
		_ = json.Unmarshal(body1, &responseBody1)

		// Verify response
		statusCode, _ := utils.ValueToInt(responseBody1["status_code"])
		assert.Equalf(t, statusCode, http.StatusAccepted, "Expected: %d, but got: %d", statusCode, http.StatusAccepted)
		assert.Equalf(t, responseBody1["message"], "logout user account is successfully!", "Expected: %s, but got: %s", responseBody["message"], "logout user account is successfully!")
		assert.Equalf(t, responseBody1["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)
	})
}
