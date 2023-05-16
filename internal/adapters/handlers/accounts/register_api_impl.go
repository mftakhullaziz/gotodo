package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"gotodo/internal/utils/payload"
	"gotodo/internal/utils/response"
	"net/http"
)

type RegisterHandlerAPI struct {
	RegisterUseCase accounts.RegisterUseCase
}

func NewRegisterHandlerAPI(registerUseCase accounts.RegisterUseCase) api.RegisterHandlerAPI {
	return &RegisterHandlerAPI{RegisterUseCase: registerUseCase}
}

func (r RegisterHandlerAPI) RegisterHandler(writer http.ResponseWriter, requests *http.Request) {
	log := logger.LoggerParent()

	registerRequest := request.RegisterRequest{}
	payload.ReadFromRequestBody(requests, &registerRequest)
	log.Info("account request: ", registerRequest)

	registerHandler, err := r.RegisterUseCase.CreateAccountUseCase(requests.Context(), registerRequest)
	errs.LoggerIfErrorWithCustomMessage(err, log, "account already created please using another 'email'")

	responses := response.CreateResponses(registerHandler, http.StatusCreated,
		"create account successfully!",
		"create account not success please check your username or email or password again!",
	)

	payload.WriteToResponseBody(writer, &responses)
}

func (r RegisterHandlerAPI) ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request) {
	//TODO implement me
	panic("implement me")
}
