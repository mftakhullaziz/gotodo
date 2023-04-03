package tasks

import (
	"gotodo/internal/domain/models/request"
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

func (t TaskHandlerAPI) CreateTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	log := helpers.LoggerParent()

	taskRequest := request.TaskRequest{}
	helpers.ReadFromRequestBody(requests, &taskRequest)
	log.Info("Task Request: ", taskRequest)

	createHandler, err := t.TaskUseCase.CreateTaskUseCase(requests.Context(), taskRequest)
	helpers.PanicIfError(err)

	responses := helpers.CreateResponses(
		createHandler,
		http.StatusCreated,
		"Create task successfully",
		"Create task not success!")

	helpers.WriteToResponseBody(writer, &responses)
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
