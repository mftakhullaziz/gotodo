package login

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotodo/config"
	"gotodo/config/database"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	accountsService "gotodo/internal/adapters/services/accounts"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	"gotodo/internal/persistence/record"
	"gotodo/internal/utils"
	"gotodo/router"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Mock function insert data account user
func mockInsertUserData(db *gorm.DB) {
	// Perform mock data insertion for account table
	mockAccount := &record.AccountRecord{
		// Set the desired mock account data
		AccountID: 1,
		UserID:    1,
		Username:  "@johndoe_test",
		Password:  utils.HashPasswordAndSalt([]byte("password")),
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_ = db.Create(mockAccount)

	// Perform mock data insertion for user detail table
	mockUserDetail := &record.UserDetailRecord{
		// Set the desired mock user detail data
		UserID:      uint(mockAccount.UserID),
		Username:    mockAccount.Username,
		Password:    mockAccount.Password,
		Email:       "johndoe@mail.com",
		Name:        "John Doe",
		MobilePhone: 6282299812,
		Address:     "Jakarta",
		Status:      "active",
		CreatedAt:   mockAccount.CreatedAt,
		UpdatedAt:   mockAccount.UpdatedAt,
	}

	_ = db.Create(mockUserDetail)
}

// Mock function database sqlite
func mockDBTest() *gorm.DB {
	// Mock function environment testing
	config.EnvironmentTest()

	// Create an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Run database migrations or any other initialization steps
	// Do migration database
	_ = database.MigrateDatabase(
		db,
		&record.TaskRecord{},
		&record.AccountRecord{},
		&record.UserDetailRecord{},
		&record.AccountLoginHistoriesRecord{},
	)

	mockInsertUserData(db)

	return db
}

// Mock function router
func mockRouter(db *gorm.DB) *httprouter.Router {
	// Initiate validator
	validate := validator.New()
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := NewLoginHandlers(loginUsecase)

	r := router.NewRouter(&loginHandler, nil, nil, nil)
	return r
}

func TestHandlers_LoginHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Login Test Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username" : "@johndoe_test", "password": "password"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, httpRequest)

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
		assert.Equalf(t, username, "@johndoe_test", "Expected: %s, but got: %s", username, "@johndoe_test")
		assert.Equalf(t, int(id), 1, "Expected: %d, but got: %d", int(id), 1)
	})

	t.Run("Login Test Is Failed", func(t *testing.T) {
		// Scenario failed is username or password is not correct
		requestBody := strings.NewReader(`{"username" : "@johndoe_test", "password": "failed"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, httpRequest)

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

func TestHandlers_LogoutHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Logout Test From User Login Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username" : "@johndoe_test", "password": "password"}`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/login", requestBody)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, httpRequest)

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
		r.ServeHTTP(recorder, httpRequest)

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
