package tasks

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type TaskUseCase interface {
	CreateTaskUseCase(ctx context.Context, request request.TaskRequest, userId int) (response.TaskResponse, error)
	UpdateTaskUseCase(ctx context.Context, request request.TaskRequest, taskId int) (response.TaskResponse, error)
	FindTaskByIdUseCase(ctx context.Context, taskId int, userId int) (response.TaskResponse, error)
	FindTaskAllUseCase(ctx context.Context, userId int) ([]response.TaskResponse, error)
	DeleteTaskUseCase(ctx context.Context, taskId int, userId int) ([]response.TaskResponse, error)
	UpdateTaskStatusUseCase(ctx context.Context, taskId int, userId int) ([]response.TaskResponse, error)
}
