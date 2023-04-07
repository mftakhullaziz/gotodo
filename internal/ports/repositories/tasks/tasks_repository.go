package tasks

import (
	"context"
	"gotodo/internal/persistence/record"
)

type TaskRecordRepository interface {
	SaveTask(ctx context.Context, taskRecord record.TaskRecord) (record.TaskRecord, error)
	FindTaskById(ctx context.Context, id int64, userId int64) (record.TaskRecord, error)
	UpdateTask(ctx context.Context, id int64, taskRecord record.TaskRecord) (record.TaskRecord, error)
	DeleteTaskById(ctx context.Context, id int64) error
	FindTaskAll(ctx context.Context, userId int64) ([]record.TaskRecord, error)
	UpdateTaskStatus(ctx context.Context, taskId int64, userId int) (record.TaskRecord, error)
}
