package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
	"net/http"
	"strconv"
)

type LoginHandlerAPI struct {
	LoginUsecase accounts.LoginUsecase
}

func NewLoginHandlerAPI(loginUsecase accounts.LoginUsecase) api.LoginHandlerAPI {
	return &LoginHandlerAPI{LoginUsecase: loginUsecase}
}

func (l LoginHandlerAPI) LoginHandler(writer http.ResponseWriter, requests *http.Request) {
	log := utils.LoggerParent()

	loginRequest := request.LoginRequest{}
	utils.ReadFromRequestBody(requests, &loginRequest)

	loginHandler, errLogin := l.LoginUsecase.LoginAccountUsecase(requests.Context(), loginRequest)
	utils.LoggerIfErrorWithCustomMessage(
		errLogin, log, "user login not success please check username or password!")

	messageIsSuccess := "login account successfully!"
	messageNotSuccess := "username and password not valid please check again!"
	result := utils.CreateResponses(
		loginHandler, http.StatusCreated, messageIsSuccess, messageNotSuccess)

	utils.WriteToResponseBody(writer, &result)
}

func (l LoginHandlerAPI) LogoutHandler(writer http.ResponseWriter, requests *http.Request) {
	token := requests.Header.Get("Authorization")
	userId, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do convert string authorized to integer
	authorizedUser, err := strconv.Atoi(userId)
	utils.LoggerIfError(err)

	// Update logout at time
	logoutHandler := l.LoginUsecase.LogoutAccountUsecase(requests.Context(), authorizedUser, token)
	utils.LoggerIfError(logoutHandler)

	// If update logout at time success then remove authorization and logout
	// Delete the Authorization header from the user's requests
	requests.Header.Del("Authorization")

	result := response.DefaultServiceResponse{
		StatusCode: http.StatusAccepted,
		Message:    "logout user account is successfully!",
		IsSuccess:  true,
		Data:       nil,
	}

	utils.WriteToResponseBody(writer, result)
}
