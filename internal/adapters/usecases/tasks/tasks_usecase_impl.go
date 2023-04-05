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

func (t TaskUseCaseImpl) FindTaskByIdUseCase(ctx context.Context, idTask int) (response.TaskResponse, error) {
	err := t.Validate.StructPartial(idTask)
	helpers.LoggerIfError(err)

	findTaskUsecase, errUsecase := t.TaskService.FindTaskByIdService(ctx, idTask)
	helpers.LoggerIfError(errUsecase)

	findTaskResponse := response.TaskResponse{
		ID:          findTaskUsecase.ID,
		UserID:      findTaskUsecase.UserID,
		Title:       findTaskUsecase.Title,
		Description: findTaskUsecase.Description,
		Completed:   findTaskUsecase.Completed,
		CompletedAt: findTaskUsecase.CompletedAt,
		CreatedAt:   findTaskUsecase.CreatedAt,
	}

	return findTaskResponse, nil
}

func (t TaskUseCaseImpl) FindTaskAllUseCase(ctx context.Context) ([]response.TaskResponse, error) {
	var findAllTaskResponse []response.TaskResponse

	findAllTaskUsecase, errUsecase := t.TaskService.FindTaskAllService(ctx)
	helpers.LoggerIfError(errUsecase)

	for _, task := range findAllTaskUsecase {
		responses := response.TaskResponse{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CompletedAt: task.CompletedAt,
			UpdatedAt:   task.UpdatedAt,
			CreatedAt:   task.CreatedAt,
		}

		findAllTaskResponse = append(findAllTaskResponse, responses)
	}

	return findAllTaskResponse, nil
}
