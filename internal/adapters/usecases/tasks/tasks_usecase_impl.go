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

	result := response.TaskResponse{
		ID:          createTaskUsecase.ID,
		UserID:      createTaskUsecase.UserID,
		Title:       createTaskUsecase.Title,
		Description: createTaskUsecase.Description,
		Completed:   createTaskUsecase.Completed,
		CompletedAt: createTaskUsecase.CompletedAt,
		CreatedAt:   createTaskUsecase.CreatedAt,
	}

	return result, nil
}

func (t TaskUseCaseImpl) UpdateTaskUseCase(ctx context.Context, request request.TaskRequest, idTask int) (response.TaskResponse, error) {
	err := t.Validate.Struct(request)
	helpers.LoggerIfError(err)

	updateTaskUsecase, errUsecase := t.TaskService.UpdateTaskService(ctx, idTask, request)
	helpers.LoggerIfError(errUsecase)

	updateTaskResult := response.TaskResponse{
		ID:          updateTaskUsecase.ID,
		UserID:      updateTaskUsecase.UserID,
		Title:       updateTaskUsecase.Title,
		Description: updateTaskUsecase.Description,
		Completed:   updateTaskUsecase.Completed,
		CompletedAt: updateTaskUsecase.CompletedAt,
		CreatedAt:   updateTaskUsecase.CreatedAt,
	}

	return updateTaskResult, nil
}
