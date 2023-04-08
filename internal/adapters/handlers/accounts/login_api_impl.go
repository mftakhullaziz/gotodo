package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
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
	log := helpers.LoggerParent()

	loginRequest := request.LoginRequest{}
	helpers.ReadFromRequestBody(requests, &loginRequest)

	loginHandler, errLogin := l.LoginUsecase.LoginAccountUsecase(requests.Context(), loginRequest)
	helpers.LoggerIfErrorWithCustomMessage(
		errLogin, log, "user login not success please check username or password!")

	messageIsSuccess := "login account successfully!"
	messageNotSuccess := "username and password not valid please check again!"
	responses := helpers.CreateResponses(
		loginHandler, http.StatusCreated, messageIsSuccess, messageNotSuccess)

	helpers.WriteToResponseBody(writer, &responses)
}

func (l LoginHandlerAPI) LogoutHandler(writer http.ResponseWriter, requests *http.Request) {
	token := requests.Header.Get("Authorization")
	userId, err := middleware.AuthenticateUser(token)
	helpers.LoggerIfError(err)

	// Do convert string authorized to integer
	authorizedUser, err := strconv.Atoi(userId)
	helpers.LoggerIfError(err)

	// Update logout at time
	logoutHandler := l.LoginUsecase.LogoutAccountUsecase(requests.Context(), authorizedUser, token)
	helpers.LoggerIfError(logoutHandler)

	// If update logout at time success then remove authorization and logout
	// Delete the Authorization header from the user's requests
	requests.Header.Del("Authorization")

	res := response.DefaultServiceResponse{
		StatusCode: http.StatusAccepted,
		Message:    "logout user account is successfully!",
		IsSuccess:  true,
		Data:       nil,
	}

	helpers.WriteToResponseBody(writer, res)
}
