package tasks

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type TaskUseCase interface {
	CreateAccountUseCase(ctx context.Context, request request.TaskRequest) (response.TaskResponse, error)
}
