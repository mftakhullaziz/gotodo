package tasks

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/ports/repositories/tasks"
	service "gotodo/internal/ports/services/tasks"
	"gotodo/internal/utils/converter"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	validate "gotodo/internal/utils/validator"
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
	// log := utils.LoggerParent()
	err := t.Validate.Struct(request)
	errs.LoggerIfError(err)

	createTaskDTO := dto.TasksDTO{
		UserID:      authorizedId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
		TaskStatus:  "active",
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	taskRecord := converter.ConvertTaskDtoToTaskRecord(createTaskDTO)
	createTaskService, err := t.TaskRepository.SaveTask(ctx, taskRecord)
	errs.LoggerIfError(err)

	taskDtoResponse := converter.ConvertTaskRecordToTaskDto(createTaskService)
	return taskDtoResponse, nil
}

func (t TaskServiceImpl) UpdateTaskService(ctx context.Context, taskId int, request request.TaskRequest) (dto.TasksDTO, error) {
	err := t.Validate.Struct(request)
	errs.LoggerIfError(err)

	// Update record Title, Description, UpdateAt
	updateTask := dto.TasksDTO{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
		UpdatedAt:   time.Time{},
	}

	taskRecord := converter.ConvertTaskDtoToTaskRecord(updateTask)
	updateService, err := t.TaskRepository.UpdateTask(ctx, int64(taskId), taskRecord)
	errs.LoggerIfError(err)

	taskDtoResponse := converter.ConvertTaskRecordToTaskDto(updateService)
	return taskDtoResponse, nil
}

func (t TaskServiceImpl) FindTaskByIdService(ctx context.Context, taskId int, userId int64) (dto.TasksDTO, error) {
	// validate task id and user id
	errIntId := validate.ValidateIntValue(taskId, int(userId))
	errs.LoggerIfError(errIntId)

	findTaskService, err := t.TaskRepository.FindTaskById(ctx, int64(taskId), userId)
	errs.LoggerIfError(err)

	findTaskResponse := converter.ConvertTaskRecordToTaskDto(findTaskService)
	return findTaskResponse, nil
}

func (t TaskServiceImpl) FindTaskAllService(ctx context.Context, userId int) ([]dto.TasksDTO, error) {
	log := logger.LoggerParent()

	findAllTaskService, err := t.TaskRepository.FindTaskAll(ctx, int64(userId))
	errs.LoggerIfError(err)
	log.Infoln("list tasks: ", findAllTaskService)

	findAllTaskResponse := converter.TaskRecordsToTaskDTOs(findAllTaskService)
	return findAllTaskResponse, nil
}

func (t TaskServiceImpl) DeleteTaskService(ctx context.Context, taskId int, userId int) error {
	err := t.Validate.StructPartial(taskId)
	errs.LoggerIfError(err)

	deleteTask := t.TaskRepository.DeleteTaskById(ctx, int64(taskId), int64(userId))
	errs.LoggerIfError(deleteTask)

	return deleteTask
}

func (t TaskServiceImpl) UpdateTaskStatusService(ctx context.Context, taskId int, userId int, completed string) (dto.TasksDTO, error) {
	err := t.Validate.StructPartial(taskId)
	errs.LoggerIfError(err)

	updateTaskStatus, errStatus := t.TaskRepository.UpdateTaskStatus(ctx, int64(taskId), int64(userId), completed)
	errs.LoggerIfError(errStatus)

	taskDtoResponse := converter.ConvertTaskRecordToTaskDto(updateTaskStatus)
	return taskDtoResponse, nil
}
