package accounts

import (
	"context"
	"gotodo/internal/infra/persistence/record"
)

type AccountRecordRepository interface {
	SaveAccount(ctx context.Context, accountRecord *record.AccountRecord) (int64, error)
	FindAccountById(ctx context.Context, id int64) (*record.AccountRecord, error)
	UpdateAccount(ctx context.Context, accountRecord *record.AccountRecord) error
	DeleteAccountById(ctx context.Context, id int64) error
	FindAccountAll(ctx context.Context) ([]record.AccountRecord, error)
}
