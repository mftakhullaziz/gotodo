package tasks

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
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

func (t TaskUseCaseImpl) CreateTaskUseCase(ctx context.Context, request request.TaskRequest) (response.TaskResponse, error) {
	err := t.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	createTaskUsecase, err := t.TaskService.CreateTaskService(ctx, request)
	if err != nil {
		panic(err)
	}

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
