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
	if errLogin != nil {
		log.Info("User login not success please check username or password!")
	}

	responses := helpers.CreateResponses(
		loginHandler,
		http.StatusCreated,
		"Login account successfully",
		"Username and password not valid please check again!")

	helpers.WriteToResponseBody(writer, &responses)
}

func (l LoginHandlerAPI) LogoutHandler(writer http.ResponseWriter, requests *http.Request) {
	panic("Implement me")
}
