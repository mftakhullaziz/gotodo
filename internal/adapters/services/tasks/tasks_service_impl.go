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
	// log := helpers.LoggerParent()
	err := t.Validate.Struct(request)
	helpers.PanicIfError(err)

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

	taskRecord := helpers.ConvertTaskDtoToTaskRecord(createTaskDTO)
	createTaskService, err := t.TaskRepository.SaveTask(ctx, taskRecord)
	helpers.LoggerIfError(err)

	taskDtoResponse := helpers.ConvertTaskRecordToTaskDto(createTaskService)
	return taskDtoResponse, nil
}

func (t TaskServiceImpl) UpdateTaskService(ctx context.Context, taskId int, request request.TaskRequest) (dto.TasksDTO, error) {
	err := t.Validate.Struct(request)
	helpers.LoggerIfError(err)

	// Update record Title, Description, UpdateAt
	updateTask := dto.TasksDTO{
		Title:       request.Title,
		Description: request.Description,
		Completed:   false,
		UpdatedAt:   time.Time{},
	}

	taskRecord := helpers.ConvertTaskDtoToTaskRecord(updateTask)
	updateService, err := t.TaskRepository.UpdateTask(ctx, int64(taskId), taskRecord)
	helpers.LoggerIfError(err)

	taskDtoResponse := helpers.ConvertTaskRecordToTaskDto(updateService)
	return taskDtoResponse, nil
}

func (t TaskServiceImpl) FindTaskByIdService(ctx context.Context, taskId int, userId int64) (dto.TasksDTO, error) {
	err := t.Validate.StructPartial(taskId)
	helpers.LoggerIfError(err)

	findTaskService, err := t.TaskRepository.FindTaskById(ctx, int64(taskId), userId)
	helpers.LoggerIfError(err)

	findTaskResponse := helpers.ConvertTaskRecordToTaskDto(findTaskService)
	return findTaskResponse, nil
}

func (t TaskServiceImpl) FindTaskAllService(ctx context.Context, userId int) ([]dto.TasksDTO, error) {
	log := helpers.LoggerParent()

	findAllTaskService, err := t.TaskRepository.FindTaskAll(ctx, int64(userId))
	helpers.LoggerIfError(err)
	log.Infoln("list tasks: ", findAllTaskService)

	findAllTaskResponse := helpers.TaskRecordsToTaskDTOs(findAllTaskService)
	return findAllTaskResponse, nil
}

func (t TaskServiceImpl) DeleteTaskService(ctx context.Context, taskId int, userId int) error {
	err := t.Validate.StructPartial(taskId)
	helpers.LoggerIfError(err)

	deleteTask := t.TaskRepository.DeleteTaskById(ctx, int64(taskId), int64(userId))
	helpers.LoggerIfError(deleteTask)

	return deleteTask
}

func (t TaskServiceImpl) UpdateTaskStatusService(ctx context.Context, taskId int, userId int, completed string) (dto.TasksDTO, error) {
	err := t.Validate.StructPartial(taskId)
	helpers.LoggerIfError(err)

	updateTaskStatus, errStatus := t.TaskRepository.UpdateTaskStatus(ctx, int64(taskId), int64(userId), completed)
	helpers.LoggerIfError(errStatus)

	taskDtoResponse := helpers.ConvertTaskRecordToTaskDto(updateTaskStatus)
	return taskDtoResponse, nil
}
