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
		TaskStatus:  "on_progress",
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

func (t TaskServiceImpl) UpdateTaskService(ctx context.Context, id int, request request.TaskRequest) (dto.TasksDTO, error) {
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
	updateService, err := t.TaskRepository.UpdateTask(ctx, int64(id), taskRecord)
	helpers.LoggerIfError(err)

	taskDtoResponse := helpers.ConvertTaskRecordToTaskDto(updateService)
	return taskDtoResponse, nil
}

func (t TaskServiceImpl) FindTaskByIdService(ctx context.Context, id int, userId int64) (dto.TasksDTO, error) {
	err := t.Validate.StructPartial(id)
	helpers.LoggerIfError(err)

	findTaskService, err := t.TaskRepository.FindTaskById(ctx, int64(id), userId)
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

func (t TaskServiceImpl) DeleteTaskService(ctx context.Context, id int) error {
	err := t.Validate.StructPartial(id)
	helpers.LoggerIfError(err)

	deleteTask := t.TaskRepository.DeleteTaskById(ctx, int64(id))
	helpers.LoggerIfError(err)

	return deleteTask
}

func (t TaskServiceImpl) UpdateTaskStatusService(ctx context.Context, id int, boolean bool) (dto.TasksDTO, error) {
	//TODO implement me
	panic("implement me")
}
