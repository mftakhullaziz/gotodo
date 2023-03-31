package tasks

import (
	"gotodo/internal/infra/persistence/record"
)

type TaskRecordRepository interface {
	Save(taskRecord *record.TaskRecord) (int64, error)
	FindById(id int64) (*record.TaskRecord, error)
	Update(taskRecord *record.TaskRecord) error
	DeleteById(id int64) error
}
