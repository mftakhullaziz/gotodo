package tasks

import (
	"github.com/julienschmidt/httprouter"
	"gotodo/internal/adapters/handlers"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	"gotodo/internal/utils"
	"net/http"
	"strconv"
	"time"
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
	log := utils.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	// Do convert string authorized to integer
	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	// Do createRequest transform to request body as json
	taskRequest := request.TaskRequest{}
	utils.ReadFromRequestBody(requests, &taskRequest)
	log.Info("task request body: ", taskRequest)

	// Do get usecase createTask function with param context, updateRequest, userId
	createHandler, err := t.TaskUseCase.CreateTaskUseCase(requests.Context(), taskRequest, authorizedUserId)
	utils.PanicIfError(err)

	// Do build response handler
	result := utils.BuildResponseWithAuthorization(
		createHandler,
		http.StatusCreated,
		int(createHandler.TaskID),
		authorized,
		"create task successful")

	// Do build write response to response body
	utils.WriteToResponseBody(writer, &result)
}

// UpdateTaskHandler : do update task based on user authorized and idTask
// Params : http.ResponseWriter, *http.Request
func (t TaskHandlerAPI) UpdateTaskHandler(writer http.ResponseWriter, requests *http.Request, ps httprouter.Params) {
	// Define logger utils
	log := utils.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	// Define to get idTask from param
	idTaskVar := ps.ByName("task_id")
	idTask, err := strconv.Atoi(idTaskVar)
	utils.LoggerIfError(err)

	// Do updateRequest transform to request body as json
	updateRequest := request.TaskRequest{}
	utils.ReadFromRequestBody(requests, &updateRequest)
	log.Infoln("update task request: ", updateRequest)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	updateTaskHandler, err := t.TaskUseCase.UpdateTaskUseCase(requests.Context(), updateRequest, idTask)
	utils.LoggerIfError(err)

	// Do build response handler
	updateTaskResponse := utils.BuildResponseWithAuthorization(
		updateTaskHandler,
		http.StatusCreated,
		int(updateTaskHandler.TaskID),
		authorized,
		"update task successfully")

	// Do build write response to response body
	utils.WriteToResponseBody(writer, &updateTaskResponse)
}

func (t TaskHandlerAPI) FindTaskHandlerById(writer http.ResponseWriter, requests *http.Request, ps httprouter.Params) {
	// Define logger utils
	log := utils.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	// Define to get idTask from param
	idTaskVar := ps.ByName("task_id")
	idTask, err := strconv.Atoi(idTaskVar)
	utils.LoggerIfError(err)
	log.Infoln("find task by id_task: ", idTask)

	// Do get usecase updateTask function with param context, updateRequest, idTask
	findTaskHandler, err := t.TaskUseCase.FindTaskByIdUseCase(requests.Context(), idTask, authorizedUserId)
	utils.LoggerIfError(err)
	log.Infoln("find task handler: ", findTaskHandler)

	findTaskHandlerResponse := utils.BuildResponseWithAuthorization(
		findTaskHandler,
		http.StatusAccepted,
		int(findTaskHandler.TaskID),
		authorized,
		"request find task successful!",
	)

	utils.WriteToResponseBody(writer, findTaskHandlerResponse)
}

func (t TaskHandlerAPI) FindTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := utils.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	findAllTaskHandler, err := t.TaskUseCase.FindTaskAllUseCase(requests.Context(), authorizedUserId)
	utils.LoggerIfError(err)

	tasksSlice := []interface{}{findAllTaskHandler}
	log.Infoln("Tasks: ", tasksSlice)

	findAllTaskHandlerResponse := utils.BuildAllResponseWithAuthorization(
		tasksSlice[0],
		"request find task successful!",
		len(findAllTaskHandler),
		time.Now().Format(handlers.FormatDatetime))

	utils.WriteToResponseBody(writer, findAllTaskHandlerResponse)
}

func (t TaskHandlerAPI) DeleteTaskHandler(writer http.ResponseWriter, requests *http.Request) {
	// Define logger utils
	log := utils.LoggerParent()

	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	// Define to get idTask from param
	params := requests.URL.Query()
	taskIdParam := params.Get("task_id")
	taskId, err := strconv.Atoi(taskIdParam)
	utils.LoggerIfError(err)

	deleteHandler := t.TaskUseCase.DeleteTaskUseCase(requests.Context(), taskId, authorizedUserId)
	if deleteHandler != nil {
		log.Infoln("delete handler not successful on task_id: ", taskId)
	}

	res := response.DefaultServiceResponse{
		StatusCode: http.StatusAccepted,
		Message:    "delete task is completed",
		IsSuccess:  true,
		Data:       "delete task is success"}

	utils.WriteToResponseBody(writer, res)
}

func (t TaskHandlerAPI) UpdateTaskStatusHandler(writer http.ResponseWriter, requests *http.Request) {
	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		result := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &result)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	// Define to get idTask from param
	params := requests.URL.Query()
	completedParam := params.Get("isCompleted")
	taskIds := params.Get("taskId")
	taskIdParam, err := strconv.Atoi(taskIds)
	utils.LoggerIfError(err)

	completedCompletedHandler, err := t.TaskUseCase.UpdateTaskStatusUseCase(
		requests.Context(), taskIdParam, authorizedUserId, completedParam)
	utils.LoggerIfError(err)

	completedHandlerResponse := utils.BuildResponseWithAuthorization(
		completedCompletedHandler,
		http.StatusAccepted,
		int(completedCompletedHandler.TaskID),
		authorized,
		"update completed task successful!")

	utils.WriteToResponseBody(writer, completedHandlerResponse)
}
