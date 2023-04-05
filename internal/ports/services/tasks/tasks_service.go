package tasks

import (
	"context"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
)

type TaskService interface {
	CreateTaskService(ctx context.Context, request request.TaskRequest, authorizedId int) (dto.TasksDTO, error)
	UpdateTaskService(ctx context.Context, id int, request request.TaskRequest) (dto.TasksDTO, error)
	FindTaskByIdService(ctx context.Context, id int) (dto.TasksDTO, error)
	FindTaskAllService(ctx context.Context) ([]dto.TasksDTO, error)
	DeleteTaskService(ctx context.Context, id int) error
	UpdateTaskStatusService(ctx context.Context, id int, boolean bool) (dto.TasksDTO, error)
}
