package register

import (
	"github.com/julienschmidt/httprouter"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
	"net/http"
)

type Handlers struct {
	RegisterUseCase accounts.RegisterUseCase
}

func NewRegisterHandlers(registerUseCase accounts.RegisterUseCase) api.RegisterHandlers {
	return &Handlers{RegisterUseCase: registerUseCase}
}

func (r Handlers) RegisterHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	log := utils.LoggerParent().Log

	registerRequest := request.RegisterRequest{}
	utils.ReadFromRequestBody(requests, &registerRequest)

	registerHandler, err := r.RegisterUseCase.CreateAccountUseCase(requests.Context(), registerRequest)
	utils.LoggerIfErrorWithCustomMessage(err, log, "account already created please using another 'email'")

	responses := utils.CreateResponses(registerHandler, http.StatusCreated,
		"create account successfully!",
		"create account not success please check your username or email or password again!",
	)

	utils.WriteToResponseBody(writer, &responses)
}

func (r Handlers) ForgotPasswordHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
