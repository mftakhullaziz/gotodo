package tasks

import (
	"context"
	"gotodo/internal/infra/persistence/record"
)

type TaskRecordRepository interface {
	SaveTask(ctx context.Context, taskRecord *record.TaskRecord) (int64, error)
	FindTaskById(ctx context.Context, id int64) (*record.TaskRecord, error)
	UpdateTask(ctx context.Context, taskRecord *record.TaskRecord) error
	DeleteTaskById(ctx context.Context, id int64) error
	FindTaskAll(ctx context.Context) ([]record.TaskRecord, error)
}
