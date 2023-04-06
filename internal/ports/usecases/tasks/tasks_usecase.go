package tasks

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type TaskUseCase interface {
	CreateTaskUseCase(ctx context.Context, request request.TaskRequest, id int) (response.TaskResponse, error)
	UpdateTaskUseCase(ctx context.Context, request request.TaskRequest, idTask int) (response.TaskResponse, error)
	FindTaskByIdUseCase(ctx context.Context, idTask int) (response.TaskResponse, error)
	FindTaskAllUseCase(ctx context.Context, userId int) ([]response.TaskResponse, error)
}
