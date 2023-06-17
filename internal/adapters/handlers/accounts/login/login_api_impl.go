package login

import (
	"github.com/julienschmidt/httprouter"
	"gotodo/internal/adapters/handlers"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
	"net/http"
	"strconv"
)

type Handlers struct {
	LoginUsecase accounts.LoginUsecase
}

func NewLoginHandlers(loginUsecase accounts.LoginUsecase) api.LoginHandlers {
	return &Handlers{LoginUsecase: loginUsecase}
}

func (h Handlers) LoginHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	log := utils.LoggerParent()

	loginRequest := request.LoginRequest{}
	utils.ReadFromRequestBody(requests, &loginRequest)

	loginHandler, errLogin := h.LoginUsecase.LoginAccountUsecase(requests.Context(), loginRequest)
	utils.LoggerIfErrorWithCustomMessage(
		errLogin, log.Log, "user login not success please check username or password!")

	messageIsSuccess := "login account successfully!"
	messageNotSuccess := "username and password not valid please check again!"
	result := utils.CreateResponses(
		loginHandler, http.StatusCreated, messageIsSuccess, messageNotSuccess)

	utils.WriteToResponseBody(writer, &result)
}

func (h Handlers) LogoutHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	token := requests.Header.Get(handlers.AuthHeaderKey)
	userId, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do convert string authorized to integer
	authorizedUser, err := strconv.Atoi(userId)
	utils.LoggerIfError(err)

	// Update logout at time
	logoutHandler := h.LoginUsecase.LogoutAccountUsecase(requests.Context(), authorizedUser, token)
	utils.LoggerIfError(logoutHandler)

	// If update logout at time success then remove authorization and logout
	// Delete the Authorization header from the user's requests
	requests.Header.Del(handlers.AuthHeaderKey)

	result := response.DefaultServiceResponse{
		StatusCode: http.StatusAccepted,
		Message:    "logout user account is successfully!",
		IsSuccess:  true,
		Data:       nil,
	}

	utils.WriteToResponseBody(writer, result)
}
