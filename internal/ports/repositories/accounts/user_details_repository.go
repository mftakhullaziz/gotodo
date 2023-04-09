package accounts

import (
	"context"
	"gotodo/internal/persistence/record"
)

type UserDetailRecordRepository interface {
	SaveUser(ctx context.Context, userRecord record.UserDetailRecord) (record.UserDetailRecord, error)
	FindUserById(ctx context.Context, userId int64) (record.UserDetailRecord, error)
	UpdateUser(ctx context.Context, userId int64, userRecord record.UserDetailRecord) (record.UserDetailRecord, error)
	DeleteUserById(ctx context.Context, userId int64) error
	FindUserAll(ctx context.Context) ([]record.UserDetailRecord, error)
}
