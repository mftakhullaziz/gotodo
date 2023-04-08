package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"net/http"
)

type RegisterHandlerAPI struct {
	RegisterUseCase accounts.RegisterUseCase
}

func NewRegisterHandlerAPI(registerUseCase accounts.RegisterUseCase) api.RegisterHandlerAPI {
	return &RegisterHandlerAPI{RegisterUseCase: registerUseCase}
}

func (r RegisterHandlerAPI) RegisterHandler(writer http.ResponseWriter, requests *http.Request) {
	log := helpers.LoggerParent()

	registerRequest := request.RegisterRequest{}
	helpers.ReadFromRequestBody(requests, &registerRequest)
	log.Info("account request: ", registerRequest)

	registerHandler, err := r.RegisterUseCase.CreateAccountUseCase(requests.Context(), registerRequest)
	helpers.LoggerIfErrorWithCustomMessage(err, log, "account already created please using another 'email'")

	responses := helpers.CreateResponses(registerHandler, http.StatusCreated,
		"create account successfully!",
		"create account not success please check your username or email or password again!",
	)

	helpers.WriteToResponseBody(writer, &responses)
}

func (r RegisterHandlerAPI) ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request) {
	//TODO implement me
	panic("implement me")
}
