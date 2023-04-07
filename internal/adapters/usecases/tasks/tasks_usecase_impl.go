package tasks

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	service "gotodo/internal/ports/services/tasks"
	"gotodo/internal/ports/usecases/tasks"
)

const formatDatetime = "2006-01-02 15:04:05"

type TaskUseCaseImpl struct {
	TaskService service.TaskService
	Validate    *validator.Validate
}

func NewTaskUseCaseImpl(taskService service.TaskService, validate *validator.Validate) tasks.TaskUseCase {
	return &TaskUseCaseImpl{TaskService: taskService, Validate: validate}
}

func (t TaskUseCaseImpl) CreateTaskUseCase(ctx context.Context, request request.TaskRequest, id int) (response.TaskResponse, error) {
	err := t.Validate.Struct(request)
	helpers.PanicIfError(err)

	createTaskUsecase, err := t.TaskService.CreateTaskService(ctx, request, id)
	helpers.PanicIfError(err)

	createTaskResponse := response.TaskResponse{
		ID:          createTaskUsecase.ID,
		UserID:      createTaskUsecase.UserID,
		Title:       createTaskUsecase.Title,
		Description: createTaskUsecase.Description,
		Completed:   createTaskUsecase.Completed,
		TaskStatus:  createTaskUsecase.TaskStatus,
		CompletedAt: createTaskUsecase.CompletedAt.Format(formatDatetime),
		CreatedAt:   createTaskUsecase.CreatedAt.Format(formatDatetime),
		UpdatedAt:   createTaskUsecase.UpdatedAt.Format(formatDatetime)}

	return createTaskResponse, nil
}

func (t TaskUseCaseImpl) UpdateTaskUseCase(ctx context.Context, request request.TaskRequest, idTask int) (response.TaskResponse, error) {
	err := t.Validate.Struct(request)
	helpers.LoggerIfError(err)

	updateTaskUsecase, errUsecase := t.TaskService.UpdateTaskService(ctx, idTask, request)
	helpers.LoggerIfError(errUsecase)
	completedTime := updateTaskUsecase.CompletedAt.Format(formatDatetime)
	updateTime := updateTaskUsecase.UpdatedAt.Format(formatDatetime)

	updateTaskResult := response.TaskResponse{
		ID:          updateTaskUsecase.ID,
		UserID:      updateTaskUsecase.UserID,
		Title:       updateTaskUsecase.Title,
		Description: updateTaskUsecase.Description,
		Completed:   updateTaskUsecase.Completed,
		TaskStatus:  updateTaskUsecase.TaskStatus,
		CompletedAt: completedTime,
		CreatedAt:   updateTaskUsecase.CreatedAt.Format(formatDatetime),
		UpdatedAt:   updateTime}

	return updateTaskResult, nil
}

func (t TaskUseCaseImpl) FindTaskByIdUseCase(ctx context.Context, idTask int, userId int) (response.TaskResponse, error) {
	err := t.Validate.StructPartial(idTask)
	helpers.LoggerIfError(err)

	findTaskUsecase, errUsecase := t.TaskService.FindTaskByIdService(ctx, idTask, int64(userId))
	helpers.LoggerIfError(errUsecase)

	findTaskResponse := response.TaskResponse{
		ID:          findTaskUsecase.ID,
		UserID:      findTaskUsecase.UserID,
		Title:       findTaskUsecase.Title,
		Description: findTaskUsecase.Description,
		Completed:   findTaskUsecase.Completed,
		TaskStatus:  findTaskUsecase.TaskStatus,
		CompletedAt: findTaskUsecase.CompletedAt.Format(formatDatetime),
		CreatedAt:   findTaskUsecase.CreatedAt.Format(formatDatetime)}

	return findTaskResponse, nil
}

func (t TaskUseCaseImpl) FindTaskAllUseCase(ctx context.Context, userId int) ([]response.TaskResponse, error) {
	var findAllTaskResponse []response.TaskResponse

	findAllTaskUsecase, errUsecase := t.TaskService.FindTaskAllService(ctx, userId)
	helpers.LoggerIfError(errUsecase)

	for _, task := range findAllTaskUsecase {
		responses := response.TaskResponse{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			TaskStatus:  task.TaskStatus,
			CompletedAt: task.CompletedAt.Format(formatDatetime),
			UpdatedAt:   task.UpdatedAt.Format(formatDatetime),
			CreatedAt:   task.CreatedAt.Format(formatDatetime)}

		findAllTaskResponse = append(findAllTaskResponse, responses)
	}

	return findAllTaskResponse, nil
}
