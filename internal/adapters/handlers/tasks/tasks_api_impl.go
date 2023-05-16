package tasks

import (
	"github.com/gorilla/mux"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"gotodo/internal/utils/payload"
	responses "gotodo/internal/utils/response"
	"net/http"
	"strconv"
	"time"
)

const (
	// formatDatetime is the format string for datetime values.
	formatDatetime = "2006-01-02 15:04:05"
	// messageUserNotAuthorized is the errors message for unauthorized user.
	messageUserNotAuthorized = "user account not authorized, please login or sign up!"
	// authHeaderKey is the key for the Authorization header.
	authHeaderKey = "Authorization"
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
	// Define logger utils
	log := logger.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	// Do convert string authorized to integer
	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	// Do createRequest transform to request body as json
	taskRequest := request.TaskRequest{}
	payload.ReadFromRequestBody(requests, &taskRequest)
	log.Info("task request body: ", taskRequest)

	// Do get usecase createTask function with param context, updateRequest, userId
	createHandler, err := t.TaskUseCase.CreateTaskUseCase(requests.Context(), taskRequest, authorizedUserId)
	errs.PanicIfError(err)

	// Do build response handler
	result := responses.BuildResponseWithAuthorization(
		createHandler,
		http.StatusCreated,
		int(createHandler.TaskID),
		authorized,
		"create task successful")

	// Do build write response to response body
	payload.WriteToResponseBody(writer, &result)
}

// UpdateTaskHandler : do update task based on user authorized and idTask
// Params : http.ResponseWriter, *http.Request
func (t TaskHandlerAPI) UpdateTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := logger.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	// Define to get idTask from param
	vars := mux.Vars(requests)
	idTaskVar := vars["task_id"]
	idTask, err := strconv.Atoi(idTaskVar)
	errs.LoggerIfError(err)

	// Do updateRequest transform to request body as json
	updateRequest := request.TaskRequest{}
	payload.ReadFromRequestBody(requests, &updateRequest)
	log.Infoln("update task request: ", updateRequest)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	updateTaskHandler, err := t.TaskUseCase.UpdateTaskUseCase(requests.Context(), updateRequest, idTask)
	errs.LoggerIfError(err)

	// Do build response handler
	updateTaskResponse := responses.BuildResponseWithAuthorization(
		updateTaskHandler,
		http.StatusCreated,
		int(updateTaskHandler.TaskID),
		authorized,
		"update task successfully")

	// Do build write response to response body
	payload.WriteToResponseBody(writer, &updateTaskResponse)
}

func (t TaskHandlerAPI) FindTaskHandlerById(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := logger.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	// Define to get idTask from param
	vars := mux.Vars(requests)
	idTaskVar := vars["task_id"]
	idTask, err := strconv.Atoi(idTaskVar)
	errs.LoggerIfError(err)
	log.Infoln("find task by id_task: ", idTask)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	findTaskHandler, err := t.TaskUseCase.FindTaskByIdUseCase(requests.Context(), idTask, authorizedUserId)
	errs.LoggerIfError(err)
	log.Infoln("find task handler: ", findTaskHandler)

	findTaskHandlerResponse := responses.BuildResponseWithAuthorization(
		findTaskHandler,
		http.StatusAccepted,
		int(findTaskHandler.TaskID),
		authorized,
		"request find task successful!",
	)

	payload.WriteToResponseBody(writer, findTaskHandlerResponse)
}

func (t TaskHandlerAPI) FindTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := logger.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	findAllTaskHandler, err := t.TaskUseCase.FindTaskAllUseCase(requests.Context(), authorizedUserId)
	errs.LoggerIfError(err)

	tasksSlice := []interface{}{findAllTaskHandler}
	log.Infoln("Tasks: ", tasksSlice)

	findAllTaskHandlerResponse := responses.BuildAllResponseWithAuthorization(
		tasksSlice[0],
		"request find task successful!",
		len(findAllTaskHandler),
		time.Now().Format(formatDatetime))

	payload.WriteToResponseBody(writer, findAllTaskHandlerResponse)
}

func (t TaskHandlerAPI) DeleteTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := logger.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	// Define to get idTask from param
	params := requests.URL.Query()
	taskIdParam := params.Get("task_id")
	taskId, err := strconv.Atoi(taskIdParam)
	errs.LoggerIfError(err)

	deleteHandler := t.TaskUseCase.DeleteTaskUseCase(requests.Context(), taskId, authorizedUserId)
	if deleteHandler != nil {
		log.Infoln("delete handler not successful on task_id: ", taskId)
	}

	res := response.DefaultServiceResponse{
		StatusCode: http.StatusAccepted,
		Message:    "delete task is completed",
		IsSuccess:  true,
		Data:       "delete task is success"}

	payload.WriteToResponseBody(writer, res)
}

func (t TaskHandlerAPI) UpdateTaskStatusHandler(writer http.ResponseWriter, requests *http.Request) {
	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := responses.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	// Define to get idTask from param
	params := requests.URL.Query()
	completedParam := params.Get("isCompleted")
	taskIds := params.Get("taskId")
	taskIdParam, err := strconv.Atoi(taskIds)
	errs.LoggerIfError(err)

	completedCompletedHandler, err := t.TaskUseCase.UpdateTaskStatusUseCase(
		requests.Context(), taskIdParam, authorizedUserId, completedParam)
	errs.LoggerIfError(err)

	completedHandlerResponse := responses.BuildResponseWithAuthorization(
		completedCompletedHandler,
		http.StatusAccepted,
		int(completedCompletedHandler.TaskID),
		authorized,
		"update completed task successful!")

	payload.WriteToResponseBody(writer, completedHandlerResponse)
}
