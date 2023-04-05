package tasks

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/repositories/tasks"
	service "gotodo/internal/ports/services/tasks"
	"time"
)

type TaskServiceImpl struct {
	TaskRepository tasks.TaskRecordRepository
	Validate       *validator.Validate
}

func NewTaskServiceImpl(taskRepository tasks.TaskRecordRepository, validate *validator.Validate) service.TaskService {
	return &TaskServiceImpl{TaskRepository: taskRepository, Validate: validate}
}

func (t TaskServiceImpl) CreateTaskService(ctx context.Context, request request.TaskRequest, authorizedId int) (dto.TasksDTO, error) {
	err := t.Validate.Struct(request)
	helpers.PanicIfError(err)

	createTask := dto.TasksDTO{
		UserID:      authorizedId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
		CompletedAt: time.Now().Add(60),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	records := helpers.TaskDTOToRecord(createTask)
	serviceTask, err := t.TaskRepository.SaveTask(ctx, records)
	if err != nil {
		panic(err)
	}
	dtoResult := helpers.TaskRecordToDTO(serviceTask)
	return dtoResult, nil
}

func (t TaskServiceImpl) UpdateTaskService(ctx context.Context, id uint8, request request.TaskRequest) (dto.TasksDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) FindTaskByIdService(ctx context.Context, id uint8) (dto.TasksDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) FindTaskAllService(ctx context.Context) ([]dto.TasksDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) DeleteTaskService(ctx context.Context, id uint8) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) UpdateTaskStatusService(ctx context.Context, id uint8, boolean bool) (dto.TasksDTO, error) {
	//TODO implement me
	panic("implement me")
}
