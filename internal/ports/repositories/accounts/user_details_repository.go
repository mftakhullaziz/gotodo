package accounts

import (
	"context"
	"gotodo/internal/infra/persistence/record"
)

type UserDetailRecordRepository interface {
	SaveUser(ctx context.Context, accountRecord *record.UserDetailRecord) (int64, error)
	FindUserById(ctx context.Context, id int64) (*record.UserDetailRecord, error)
	UpdateUser(ctx context.Context, accountRecord *record.UserDetailRecord) error
	DeleteUserById(ctx context.Context, id int64) error
	FindUserAll(ctx context.Context) ([]record.UserDetailRecord, error)
}
