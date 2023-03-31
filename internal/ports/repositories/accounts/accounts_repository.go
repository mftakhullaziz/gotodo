package accounts

import "gotodo/internal/infra/persistence/record"

type AccountRecordRepository interface {
	Save(taskRecord *record.AccountRecord) (int64, error)
	FindById(id int64) (*record.AccountRecord, error)
	Update(taskRecord *record.AccountRecord) error
	DeleteById(id int64) error
}
