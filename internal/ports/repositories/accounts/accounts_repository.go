package accounts

import (
	"context"
	"gotodo/internal/persistence/record"
)

type AccountRecordRepository interface {
	SaveAccount(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindAccountById(ctx context.Context, id int64) (record.AccountRecord, error)
	UpdateAccount(ctx context.Context, id int64, accountRecord record.AccountRecord) (record.AccountRecord, error)
	DeleteAccountById(ctx context.Context, id int64) error
	FindAccountAll(ctx context.Context) ([]record.AccountRecord, error)
	FindAccountByEmail(ctx context.Context, email string) bool
}
