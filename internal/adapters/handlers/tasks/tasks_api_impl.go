package tasks

import (
	"github.com/julienschmidt/httprouter"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	"net/http"
)

type TaskHandlerAPI struct {
	TaskUseCase tasks.TaskUseCase
}

func NewTaskHandlerAPI(taskUseCase tasks.TaskUseCase) api.TaskHandlerAPI {
	return &TaskHandlerAPI{TaskUseCase: taskUseCase}
}

func (t TaskHandlerAPI) CreateTaskHandler(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	taskRequest := request.TaskRequest{}
	helpers.ReadFromRequestBody(requests, &taskRequest)

	createHandler, err := t.TaskUseCase.CreateAccountUseCase(requests.Context(), taskRequest)
	helpers.PanicIfError(err)
	rest := response.DefaultServiceResponse{
		StatusCode: 201,
		Message:    "Create Task Success",
		IsSuccess:  true,
		Data:       createHandler,
	}
	helpers.WriteToResponseBody(writer, &rest)
}
