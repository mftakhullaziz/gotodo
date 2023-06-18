package tasks

import (
	"context"
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
	tasksRepository "gotodo/internal/adapters/repositories/tasks"
	accountsService "gotodo/internal/adapters/services/accounts"
	tasksService "gotodo/internal/adapters/services/tasks"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	tasksUsecase "gotodo/internal/adapters/usecases/tasks"
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

	mockTask1 := &record.TaskRecord{
		TaskID:      1,
		UserID:      mockAccount.UserID,
		Title:       "Create first tasks to create fixing bugs",
		Description: "Completed tasks to finish fixing bugs hotfix in production",
		Completed:   false,
		TaskStatus:  "active",
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	_ = db.Create(mockTask1)

	mockTask2 := &record.TaskRecord{
		TaskID:      2,
		UserID:      mockAccount.UserID,
		Title:       "Building api list",
		Description: "Build api all in service consumer",
		Completed:   false,
		TaskStatus:  "active",
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	_ = db.Create(mockTask2)
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
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(db, validate)
	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := login.NewLoginHandlers(loginUsecase)

	// Task Handler
	taskRepository := tasksRepository.NewTaskRepositoryImpl(db, validate)
	taskService := tasksService.NewTaskServiceImpl(taskRepository, validate)
	taskUsecase := tasksUsecase.NewTaskUseCaseImpl(taskService, validate)
	taskHandler := NewTaskHandlers(taskUsecase)

	r := router.NewRouter(&loginHandler, nil, nil, &taskHandler)
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

func TestHandlers_CreateTaskHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Test Create Task Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`
			{
				"title": "Create first tasks to create fixing bugs",
				"description": "Completed tasks to finish fixing bugs hotfix in production"
			}
		`)

		httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/task/create", requestBody)

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
		assert.Equalf(t, responseBody["message"], "create task successfully!", "Expected: %s, but got: %s", responseBody["message"], "create task successfully!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}

func TestHandlers_UpdateTaskHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Test Update Task Is Success", func(t *testing.T) {
		requestBody := strings.NewReader(`
			{
				"title": "Create first tasks to create fixing bugs in monday",
				"description": "Completed tasks to finish fixing bugs hotfix in production in monday"
			}
		`)

		// Set the task ID parameter value
		taskId := "1"
		url := "http://localhost:3000/api/v1/task/update/" + taskId

		httpRequest := httptest.NewRequest(http.MethodPut, url, requestBody)

		// Create a new httprouter.Params object and add the task ID parameter to it
		params := httprouter.Params{
			httprouter.Param{
				Key:   "task_id",
				Value: taskId,
			},
		}

		// Set the params object in the request context
		ctx := context.WithValue(httpRequest.Context(), httprouter.ParamsKey, params)
		httpRequest = httpRequest.WithContext(ctx)

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
		assert.Equalf(t, responseBody["message"], "update task successfully", "Expected: %s, but got: %s", responseBody["message"], "update task successfully")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}

func TestHandlers_UpdateTaskStatusHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Test Update Task Status", func(t *testing.T) {
		// Build url
		taskID := "1"
		isCompleted := "true"
		url := "http://localhost:3000/api/v1/task/update_status?taskId=" + taskID + "&isCompleted=" + isCompleted

		httpRequest := httptest.NewRequest(http.MethodPut, url, nil)

		// Add URL parameters to the request
		params := httprouter.Params{
			httprouter.Param{Key: "taskId", Value: taskID},
			httprouter.Param{Key: "isCompleted", Value: isCompleted},
		}

		// Set the params object in the request context
		ctx := context.WithValue(httpRequest.Context(), httprouter.ParamsKey, params)
		httpRequest = httpRequest.WithContext(ctx)

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
		assert.Equalf(t, statusCode, http.StatusAccepted, "Expected: %d, but got: %d", statusCode, http.StatusAccepted)
		assert.Equalf(t, responseBody["message"], "update completed task successful!", "Expected: %s, but got: %s", responseBody["message"], "update completed task successful!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}

func TestHandlers_FindTaskHandler(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Test Find Task All", func(t *testing.T) {
		httpRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/task/find", nil)

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
		assert.Equalf(t, responseBody["message"], "request find task successful!", "Expected: %s, but got: %s", responseBody["message"], "request find task successful!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])

		totalData, _ := utils.ValueToInt(responseBody["total_data"])
		assert.GreaterOrEqualf(t, totalData, 0, "Expected: %d, but got: %d", totalData, 0)
	})
}

func TestHandlers_FindTaskHandlerById(t *testing.T) {
	db := mockDBTest()
	r := mockRouter(db)

	t.Run("Test Find Task Handler By Id", func(t *testing.T) {
		// Set the task ID parameter value
		taskId := "1"
		url := "http://localhost:3000/api/v1/task/find/" + taskId

		httpRequest := httptest.NewRequest(http.MethodGet, url, nil)

		// Create a new httprouter.Params object and add the task ID parameter to it
		params := httprouter.Params{
			httprouter.Param{
				Key:   "task_id",
				Value: taskId,
			},
		}

		// Set the params object in the request context
		ctx := context.WithValue(httpRequest.Context(), httprouter.ParamsKey, params)
		httpRequest = httpRequest.WithContext(ctx)

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
		assert.Equalf(t, statusCode, http.StatusAccepted, "Expected: %d, but got: %d", statusCode, http.StatusAccepted)
		assert.Equalf(t, responseBody["message"], "request find task successful!", "Expected: %s, but got: %s", responseBody["message"], "request find task successful!")
		assert.Equalf(t, responseBody["is_success"], true, "Expected: %s, but got: %s", responseBody["is_success"], true)

		// Verify response data
		assert.NotEmptyf(t, responseBody["data"], "Expected: %s", responseBody["data"])
		assert.NotNilf(t, responseBody["data"], "Expected: %s", responseBody["data"])
	})
}

func TestHandlers_DeleteTaskHandler(t *testing.T) {
}
