package accounts

import (
	request "gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
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
	log.Info("Account request: ", registerRequest)

	registerHandler, err := r.RegisterUseCase.CreateAccountUseCase(requests.Context(), registerRequest)
	helpers.PanicIfError(err)

	accountResponse := response.DefaultServiceResponse{
		StatusCode: 201,
		Message:    "Create Account Success",
		IsSuccess:  true,
		Data:       registerHandler,
	}
	log.Info("Account created successfully")
	helpers.WriteToResponseBody(writer, &accountResponse)
}

func (r RegisterHandlerAPI) ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request) {
	//TODO implement me
	panic("implement me")
}
