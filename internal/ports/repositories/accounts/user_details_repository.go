package accounts

import "gotodo/internal/infra/persistence/record"

type UserDetailRecordRepository interface {
	SaveUser(taskRecord *record.UserDetailRecord) (int64, error)
	FindUserById(id int64) (*record.UserDetailRecord, error)
	UpdateUser(taskRecord *record.UserDetailRecord) error
	DeleteUserById(id int64) error
	FindUserAll(taskRecord record.UserDetailRecord) error
}
