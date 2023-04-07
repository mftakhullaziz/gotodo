package tasks

import (
	"context"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
)

type TaskService interface {
	CreateTaskService(ctx context.Context, request request.TaskRequest, authorizedId int) (dto.TasksDTO, error)
	UpdateTaskService(ctx context.Context, taskId int, request request.TaskRequest) (dto.TasksDTO, error)
	FindTaskByIdService(ctx context.Context, taskId int, userId int64) (dto.TasksDTO, error)
	FindTaskAllService(ctx context.Context, userId int) ([]dto.TasksDTO, error)
	DeleteTaskService(ctx context.Context, taskId int, userId int) error
	UpdateTaskStatusService(ctx context.Context, taskId int, userId int, completed string) (dto.TasksDTO, error)
}
