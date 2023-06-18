package accounts

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotodo/config"
	"gotodo/config/database"
	"gotodo/internal/adapters/handlers/accounts/login"
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
	// Repo init
	userRepository := accountsRepository.NewUserDetailRepositoryImpl(db, validate)
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)

	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := login.NewLoginHandlers(loginUsecase)

	// User Detail Handler
	userDetailService := accountsService.NewUserDetailServiceImpl(userRepository, accountRepository, validate)
	userDetailUsecase := accountsUsecase.NewUserDetailUsecaseImpl(userDetailService, validate)
	userDetailHandler := NewUserDetailHandlers(userDetailUsecase)

	r := router.NewRouter(&loginHandler, nil, &userDetailHandler, nil)
	return r
}

// Mock Login to get Authorization Key
func mockLogin() string {
	db := mockDBTest()
	r := mockRouter(db)

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

	return authorizationKey
}

func TestHandlers_FindDataUserDetailHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Find Data User Is Success", func(t *testing.T) {
		httpRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/user/find", nil)

		// Mock Authorization
		authorizationKey := mockLogin()

		// Setup Header key for parsing token in authorization
		headers := make(http.Header)
		headers.Set("Content-Type", "application/json")
		headers.Set("Authorization", authorizationKey)
		httpRequest.Header = headers

		// Execute the request
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, httpRequest)

		data := recorder.Result()
		body, _ := io.ReadAll(data.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		// Verify response
		statusCode, _ := utils.ValueToInt(responseBody["status_code"])
		assert.Equalf(t, statusCode, http.StatusOK, "Expected: %d, but got: %d", statusCode, http.StatusOK)
		assert.Equalf(t, responseBody["message"], "find user detail successfully!", "Expected: %s, but got: %s", responseBody["message"], "find user detail successfully!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}

func TestHandlers_UpdateUserDetailHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Update User Detail Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`
			{
				"email": "",
				"password": "password1",
				"name": "",
				"mobilePhone": 6282244444,
				"address": "Bandar lampung",
				"status": "active"
			}`,
		)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/user/edit", requestBody)

		// Mock Authorization
		authorizationKey := mockLogin()

		// Setup Header key for parsing token in authorization
		headers := make(http.Header)
		headers.Set("Content-Type", "application/json")
		headers.Set("Authorization", authorizationKey)
		httpRequest.Header = headers

		// Execute the request
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, httpRequest)

		data := recorder.Result()
		body, _ := io.ReadAll(data.Body)
		var responseBody map[string]interface{}
		_ = json.Unmarshal(body, &responseBody)

		// Verify response
		statusCode, _ := utils.ValueToInt(responseBody["status_code"])
		assert.Equalf(t, statusCode, http.StatusCreated, "Expected: %d, but got: %d", statusCode, http.StatusCreated)
		assert.Equalf(t, responseBody["message"], "update user successfully!", "Expected: %s, but got: %s", responseBody["message"], "update user successfully!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}
