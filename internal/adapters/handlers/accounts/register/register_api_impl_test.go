package register

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
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
)

// Mock DB In SQLite
func mockDbTest() *gorm.DB {
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

	return db
}

func mockRouter(db *gorm.DB) *httprouter.Router {
	// Initiate validator
	validate := validator.New()
	userRepository := accountsRepository.NewUserDetailRepositoryImpl(db, validate)
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	accountService := accountsService.NewRegisterServiceImpl(accountRepository, userRepository, validate)
	accountUsecase := accountsUsecase.NewRegisterUseCaseImpl(accountService, validate)
	accountHandler := NewRegisterHandlers(accountUsecase)

	r := router.NewRouter(nil, &accountHandler, nil, nil)
	return r
}

func TestHandlers_RegisterHandler(t *testing.T) {
	db := mockDbTest()
	r := mockRouter(db)

	t.Run("Register New User Test Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username": "@johndoe_test", "email": "johndoetest04@mail.com", "password": "password" }`)
		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/authenticate/register", requestBody)

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
		assert.Equalf(t, responseBody["message"], "create account successfully!", "Expected: %s, but got: %s", responseBody["message"], "create account successfully!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})

}
