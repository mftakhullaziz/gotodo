package accounts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockLoginUsecase struct {
	loginResponse response.LoginResponse
	loginError    error
	logoutError   error
}

func (m *mockLoginUsecase) LoginAccountUsecase(_ context.Context, _ request.LoginRequest) (response.LoginResponse, error) {
	return m.loginResponse, m.loginError
}

func (m *mockLoginUsecase) LogoutAccountUsecase(ctx context.Context, userID int, token string) error {
	return m.logoutError
}

func TestLoginHandlerAPI_LoginHandler(t *testing.T) {
	// Create a mock LoginUsecase
	loginUsecase := &mockLoginUsecase{
		loginResponse: response.LoginResponse{
			AccountID: 1,
			Username:  "test_user",
			Password:  "test_pass",
		},
		loginError: nil,
	}

	// Create a new LoginHandlerAPI instance
	loginHandler := NewLoginHandlerAPI(loginUsecase)

	// Create a test request with a JSON payload
	loginReq := request.LoginRequest{
		Username: "test_user",
		Password: "test_pass",
	}
	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the LoginHandler method
	loginHandler.LoginHandler(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
	fmt.Println(recorder.Body)

	// Parse the response body
	var responseBody response.DefaultServiceResponse
	err := json.NewDecoder(recorder.Body).Decode(&responseBody)
	assert.NoError(t, err)

	// Verify the response
	assert.True(t, responseBody.IsSuccess)
	assert.Equal(t, "login account successfully!", responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	fmt.Println(responseBody.Data)

	loginResponse, ok := responseBody.Data.(response.LoginResponse)

	fmt.Println(loginResponse.AccountID)

	assert.True(t, ok, "response.Data is not of type response.LoginResponse")

	//
	//assert.Equal(t, 1, loginResponse.AccountID)
	//assert.Equal(t, "testuser", loginResponse.Username)
	//assert.Equal(t, "testtoken", loginResponse.Token)
}
