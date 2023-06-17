package login

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"gotodo/config"
	"gotodo/config/database"
	accountsRepository "gotodo/internal/adapters/repositories/accounts"
	accountsService "gotodo/internal/adapters/services/accounts"
	accountsUsecase "gotodo/internal/adapters/usecases/accounts"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Initiate Setup Test Database
func SetupTestDB() *gorm.DB {
	// Initiate context
	ctx := context.Background()
	env := config.LoadEnv(".env")
	// Do function database connection
	db, _ := database.NewDatabaseConnection(ctx, env)
	return db
}

func SetupRouter() func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Initiate validator
	validate := validator.New()
	accountRepository := accountsRepository.NewAccountsRepositoryImpl(SetupTestDB(), validate)

	loginService := accountsService.NewLoginServiceImpl(accountRepository, validate)
	loginUsecase := accountsUsecase.NewLoginUsecaseImpl(loginService, validate)
	loginHandler := NewLoginHandlerAPI(loginUsecase)

	return loginHandler.LoginHandler
}

type MockLoginUsecase struct {
	LoginResponse response.LoginResponse
	Err           error
}

func (m *MockLoginUsecase) LoginAccountUsecase(_ context.Context, _ request.LoginRequest) (response.LoginResponse, error) {
	return m.LoginResponse, m.Err
}

func (m *MockLoginUsecase) LogoutAccountUsecase(_ context.Context, _ int, _ string) error {
	return m.Err
}

func TestLoginHandlerAPI_LoginHandler(t *testing.T) {
	// Login is success
	t.Run("Login is success", func(t *testing.T) {
		// Create a mock LoginUsecase
		mockUsecase := &MockLoginUsecase{
			LoginResponse: response.LoginResponse{
				AccountID:         int(1),
				Username:          "johndoe",
				Password:          utils.HashPasswordAndSalt([]byte("password")),
				LoginCreationTime: time.Now(),
				LoginToken:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODY5ODUyODJ9.iSg_zBE5oeBywFoDb8hItZe-agiQLKTuUYGQRj307P8",
			},
			Err: nil,
		}

		// Create a new LoginHandlerAPI instance
		handler := NewLoginHandlerAPI(mockUsecase).LoginHandler

		router := httprouter.New()
		router.POST("/login", handler)

		payload := request.LoginRequest{Username: "johndoe", Password: "password"}
		jsonMarshal, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonMarshal))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)
	})
}

//func TestLoginHandlerAPI_LoginHandler(t *testing.T) {
//	// Login is success
//	t.Run("Login is success", func(t *testing.T) {
//		// Create a mock LoginUsecase
//		var accountId int
//		accountId = 1
//		mockUsecase := &mockLoginUsecase{
//			response: response.LoginResponse{
//				AccountID: accountId,
//				Username:  "johndoe",
//				Password:  utils.HashPasswordAndSalt([]byte("password")),
//				LoginAt:   time.Now(),
//				Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODY5ODUyODJ9.iSg_zBE5oeBywFoDb8hItZe-agiQLKTuUYGQRj307P8",
//			},
//			err: nil,
//		}
//
//		// Create a new LoginHandlerAPI instance
//		handler := NewLoginHandlerAPI(mockUsecase)
//
//		// Create a test request with a JSON payload
//		payload := request.LoginRequest{Username: "johndoe", Password: "password"}
//		requestBody, _ := json.Marshal(payload)
//		requestUrl, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
//
//		// Create a test response recorder
//		recorder := httptest.NewRecorder()
//
//		// Call the LoginHandler method
//		handler.LoginHandler(recorder, requestUrl)
//
//		// Parse the response body
//		var responseBody response.DefaultServiceResponse
//		_ = json.NewDecoder(recorder.Body).Decode(&responseBody)
//		// Verify response data
//		data, _ := utils.InterfaceToMap(responseBody.Data)
//		// get password from map
//		pass := fmt.Sprintf("%s", data["password"])
//		comparedPassword, _ := utils.ComparedPassword(pass, []byte(payload.Password))
//		assert.True(t, comparedPassword)
//
//		// verify response status
//		assert.Equal(t, http.StatusCreated, responseBody.StatusCode)
//
//		// Verify the response
//		assert.True(t, responseBody.IsSuccess)
//		assert.Equal(t, "login account successfully!", responseBody.Message)
//		assert.NotNil(t, responseBody.Data)
//
//		assert.Equal(t, data["account_id"], float64(1))
//		assert.Equal(t, data["username"], payload.Username)
//	})
//
//	// Login is failed
//	t.Run("Login is failed", func(t *testing.T) {
//		// Mock login usecase failed condition
//		mockUsecase := &mockLoginUsecase{
//			response: response.LoginResponse{},
//			err:      nil,
//		}
//
//		// Create a new LoginHandlerAPI instance
//		handler := NewLoginHandlerAPI(mockUsecase)
//
//		// Create a test request with a JSON payload
//		payload := request.LoginRequest{Username: "johndoe", Password: "password1"}
//		requestBody, _ := json.Marshal(payload)
//		requestUrl, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
//
//		// Create a test response recorder
//		recorder := httptest.NewRecorder()
//
//		// Call the function of LoginHandler
//		handler.LoginHandler(recorder, requestUrl)
//
//		// Parse the response body
//		var responseBody response.DefaultServiceResponse
//		_ = json.NewDecoder(recorder.Body).Decode(&responseBody)
//
//		// decode map data
//		data, _ := utils.InterfaceToMap(responseBody.Data)
//		// get password from map
//		pass := fmt.Sprintf("%s", data["password"])
//		comparedPassword, _ := utils.ComparedPassword(pass, []byte(payload.Password))
//		assert.False(t, comparedPassword)
//
//		// verify response status
//		assert.Equal(t, http.StatusBadRequest, responseBody.StatusCode)
//
//		// Verify the response
//		assert.False(t, responseBody.IsSuccess)
//		assert.Equal(t, "username and password not valid please check again!", responseBody.Message)
//		assert.Nil(t, responseBody.Data)
//	})
//
//	t.Run("test", func(t *testing.T) {
//		payload := request.LoginRequest{Username: "johndoe", Password: "password1"}
//		requestBody, _ := json.Marshal(payload)
//		requestUrl, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
//		fmt.Println(requestUrl)
//		//res, err := http.Post(requestUrl.URL.String())
//	})
//}
//
//func TestLoginHandlerAPI_LoginHandler_Success(t *testing.T) {
//	// Create a mock LoginUsecase with the desired behavior
//	mockLoginUsecase := &mockLoginUsecase{
//		LoginAccountUsecaseFunc: func(ctx context.Context, request request.LoginRequest) (*response.LoginResponse, error) {
//			// Return a successful login response
//			return &response.LoginResponse{}, nil
//		},
//	}
//
//	// Create a new instance of LoginHandlerAPI with the mockLoginUsecase
//	loginHandlerAPI := NewLoginHandlerAPI(mockLoginUsecase)
//
//	// Create a test HTTP request with the desired payload
//	requestPayload := `{"username": "test", "password": "pass"}`
//	req, err := http.NewRequest("POST", "/login", strings.NewReader(requestPayload))
//	if err != nil {
//		t.Fatalf("Failed to create request: %v", err)
//	}
//
//	// Create a ResponseRecorder to record the response
//	rr := httptest.NewRecorder()
//
//	// Call the LoginHandler method
//	loginHandlerAPI.LoginHandler(rr, req)
//
//	// Check the response status code
//	if rr.Code != http.StatusCreated {
//		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
//	}
//
//	// Perform additional assertions on the response if needed
//}
//
//func TestLoginHandlerAPI_LoginHandler_Failure(t *testing.T) {
//	// Create a mock LoginUsecase with the desired behavior
//	mockLoginUsecase := &mockLoginUsecases{
//		LoginAccountUsecaseFunc: func(ctx context.Context, request request.LoginRequest) (*response.LoginResponse, error) {
//			// Return an error indicating login failure
//			return nil, errors.New("username or password is invalid")
//		},
//	}
//
//	// Create a new instance of LoginHandlerAPI with the mockLoginUsecase
//	loginHandlerAPI := NewLoginHandlerAPI(mockLoginUsecase)
//
//	// Create a test HTTP request with the desired payload
//	requestPayload := `{"username": "test", "password": "pass"}`
//	req, err := http.NewRequest("POST", "/login", strings.NewReader(requestPayload))
//	if err != nil {
//		t.Fatalf("Failed to create request: %v", err)
//	}
//
//	// Create a ResponseRecorder to record the response
//	rr := httptest.NewRecorder()
//
//	// Call the LoginHandler method
//	loginHandlerAPI.LoginHandler(rr, req)
//
//	// Check the response status code
//	if rr.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
//	}
//
//	// Perform additional assertions on the response if needed
//}
//
//// MockLoginUsecase is a mock implementation of the accounts.LoginUsecase interface
//type mockLoginUsecases struct {
//	LoginAccountUsecaseFunc func(ctx context.Context, request request.LoginRequest) (*response.LoginResponse, error)
//}
//
//func (m *mockLoginUsecase) LoginAccountUsecases(ctx context.Context, request request.LoginRequest) (*response.LoginResponse, error) {
//	return m.LoginAccountUsecaseFunc()
//}
//
//func (m *mockLoginUsecase) LoginAccountUsecaseFunc(ctx context.Context, loginRequest request.LoginRequest) *response.LoginResponse {
//
//}

//type LoginUsecase interface {
//	LoginAccountUsecase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error)
//}
//
//type MockLoginUsecase struct {
//	LoginResponse response.LoginResponse
//	Err           error
//}
//
//func (m *MockLoginUsecase) LoginAccountUsecase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error) {
//	return m.LoginResponse, m.Err
//}
//
//type mockLoginUsecases struct {
//	response response.LoginResponse
//	err      error
//}
//
//func (m *mockLoginUsecases) LoginAccountUsecase(_ context.Context, _ request.LoginRequest) (response.LoginResponse, error) {
//	return m.response, m.err
//}

//func TestAPILogin(t *testing.T) {
//	// Create a mock LoginUsecase
//	var accountId int
//	accountId = 1
//	mockUsecase := &mockLoginUsecases{
//		response: response.LoginResponse{
//			AccountID: accountId,
//			Username:  "johndoe",
//			Password:  utils.HashPasswordAndSalt([]byte("password")),
//			LoginAt:   time.Now(),
//			Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODY5ODUyODJ9.iSg_zBE5oeBywFoDb8hItZe-agiQLKTuUYGQRj307P8",
//		},
//		err: nil,
//	}
//
//	NewLoginHandlerAPI()
//
//	handler := NewLoginHandlerAPI(mockUsecase).LoginHandler
//
//	//handlers := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	//	//controllers.NewBookController().GetBook(w, r, httprouter.Params{})
//	//	NewLoginHandlerAPI(mockUsecase).LogoutHandler(w, r, httprouter.Params{})
//	//})
//
//}
