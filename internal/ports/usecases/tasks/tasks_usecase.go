package tasks

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type TaskUseCase interface {
	CreateTaskUseCase(ctx context.Context, request request.TaskRequest) (response.TaskResponse, error)
}
