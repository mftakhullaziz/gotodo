package tasks

import (
	"github.com/sirupsen/logrus"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	"net/http"
)

var log logrus.Logger

type TaskHandlerAPI struct {
	TaskUseCase tasks.TaskUseCase
}

func NewTaskHandlerAPI(taskUseCase tasks.TaskUseCase) api.TaskHandlerAPI {
	return &TaskHandlerAPI{TaskUseCase: taskUseCase}
}

func (t TaskHandlerAPI) CreateTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	log.Debug("Received request to create a task")

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

	log.Info("Task created successfully")
	helpers.WriteToResponseBody(writer, &rest)
}

func (t TaskHandlerAPI) UpdateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t TaskHandlerAPI) FindTaskHandlerById(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t TaskHandlerAPI) FindTaskHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t TaskHandlerAPI) DeleteTaskHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
