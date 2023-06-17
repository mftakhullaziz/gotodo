package accounts

//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"gotodo/internal/domain/models/request"
//	"gotodo/internal/domain/models/response"
//	"gotodo/internal/utils"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//	"time"
//)
//
//type mockLoginUsecase struct {
//	response response.LoginResponse
//	err      error
//}
//
//func (m *mockLoginUsecase) LoginAccountUsecase(_ context.Context, _ request.LoginRequest) (response.LoginResponse, error) {
//	return m.response, m.err
//}
//
//func (m *mockLoginUsecase) LogoutAccountUsecase(ctx context.Context, userID int, token string) error {
//	return m.err
//}
//
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
