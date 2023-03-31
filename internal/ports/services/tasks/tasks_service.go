package tasks

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type TaskService interface {
	CreateTask(ctx context.Context, request request.TaskRequest) (response.TaskResponse, error)
	UpdateTask(ctx context.Context, id uint8, request request.TaskRequest) (response.TaskResponse, error)
	FindTaskById(ctx context.Context, id uint8) (response.TaskResponse, error)
	FindTaskAll(ctx context.Context) ([]response.TaskAllResponse, error)
	DeleteTask(ctx context.Context, id uint8) error
}
