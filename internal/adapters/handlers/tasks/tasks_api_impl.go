package tasks

import (
	"github.com/gorilla/mux"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	"net/http"
	"strconv"
)

type TaskHandlerAPI struct {
	TaskUseCase tasks.TaskUseCase
}

func NewTaskHandlerAPI(taskUseCase tasks.TaskUseCase) api.TaskHandlerAPI {
	return &TaskHandlerAPI{TaskUseCase: taskUseCase}
}

// CreateTaskHandler : do update task based on user authorized
// Params : http.ResponseWriter, *http.Request
func (t TaskHandlerAPI) CreateTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger helpers
	log := helpers.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get("Authorization")
	authorized, err := middleware.AuthenticateUser(token)
	helpers.LoggerIfError(err)

	// Do convert string authorized to integer
	authorizedUserId, err := strconv.Atoi(authorized)
	helpers.LoggerIfError(err)

	// Do createRequest transform to request body as json
	taskRequest := request.TaskRequest{}
	helpers.ReadFromRequestBody(requests, &taskRequest)
	log.Info("Task Request: ", taskRequest)

	// Do get usecase createTask function with param context, updateRequest, userId
	createHandler, err := t.TaskUseCase.CreateTaskUseCase(requests.Context(), taskRequest, authorizedUserId)
	helpers.PanicIfError(err)

	// Do build response handler
	responses := helpers.BuildResponseWithAuthorization(createHandler, http.StatusCreated, authorized,
		"Create task successfully", "Create task not success!")

	// Do build write response to response body
	helpers.WriteToResponseBody(writer, &responses)
}

// UpdateTaskHandler : do update task based on user authorized and idTask
// Params : http.ResponseWriter, *http.Request
func (t TaskHandlerAPI) UpdateTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger helpers
	log := helpers.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get("Authorization")
	authorizedUser, err := middleware.AuthenticateUser(token)
	helpers.LoggerIfError(err)

	// Define to get idTask from param
	vars := mux.Vars(requests)
	idTaskVar := vars["idTask"]
	idTask, err := strconv.Atoi(idTaskVar)
	helpers.LoggerIfError(err)

	// Do updateRequest transform to request body as json
	updateRequest := request.TaskRequest{}
	helpers.ReadFromRequestBody(requests, &updateRequest)
	log.Infoln("Update task request: ", updateRequest)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	updateTaskHandler, err := t.TaskUseCase.UpdateTaskUseCase(requests.Context(), updateRequest, idTask)
	helpers.LoggerIfError(err)

	// Do build response handler
	updateTaskResponse := helpers.BuildResponseWithAuthorization(updateTaskHandler, http.StatusCreated, authorizedUser,
		"Update task successfully", "Update task is fail please check your parameters!")

	// Do build write response to response body
	helpers.WriteToResponseBody(writer, &updateTaskResponse)
}

func (t TaskHandlerAPI) FindTaskHandlerById(writer http.ResponseWriter, requests *http.Request) {
	// Define logger helpers
	log := helpers.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get("Authorization")
	authorized, err := middleware.AuthenticateUser(token)
	helpers.LoggerIfError(err)

	// Define to get idTask from param
	vars := mux.Vars(requests)
	idTaskVar := vars["idTask"]
	idTask, err := strconv.Atoi(idTaskVar)
	helpers.LoggerIfError(err)
	log.Infoln("Find task by idTask: ", idTask)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	findTaskHandler, err := t.TaskUseCase.FindTaskByIdUseCase(requests.Context(), idTask)
	helpers.LoggerIfError(err)

	findTaskHandlerResponse := helpers.BuildResponseWithAuthorization(findTaskHandler, http.StatusAccepted, authorized,
		"API request find task successful", "Request is failed please check again taskId!")

	helpers.WriteToResponseBody(writer, findTaskHandlerResponse)
}

func (t TaskHandlerAPI) FindTaskHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (t TaskHandlerAPI) DeleteTaskHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
