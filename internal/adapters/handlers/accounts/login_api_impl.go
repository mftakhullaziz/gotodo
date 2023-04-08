package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"net/http"
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

	loginHandler, errLogin := l.LoginUsecase.LoginAccountUseCase(requests.Context(), loginRequest)
	helpers.LoggerIfErrorWithCustomMessage(
		errLogin, log, "user login not success please check username or password!")

	messageIsSuccess := "login account successfully!"
	messageNotSuccess := "username and password not valid please check again!"
	responses := helpers.CreateResponses(
		loginHandler, http.StatusCreated, messageIsSuccess, messageNotSuccess)

	helpers.WriteToResponseBody(writer, &responses)
}

func (l LoginHandlerAPI) LogoutHandler(writer http.ResponseWriter, requests *http.Request) {
	// Delete the Authorization header from the user's requests
	requests.Header.Del("Authorization")

}
