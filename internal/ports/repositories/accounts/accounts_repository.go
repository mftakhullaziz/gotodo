package accounts

import (
	"context"
	"gotodo/internal/persistence/record"
	"gotodo/internal/utils"
)

type AccountRecordRepository interface {
	SaveAccount(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindAccountById(ctx context.Context, id int64) (record.AccountRecord, error)
	UpdateAccount(ctx context.Context, id int64, accountRecord record.AccountRecord) (record.AccountRecord, error)
	DeleteAccountById(ctx context.Context, id int64) error
	FindAccountAll(ctx context.Context) ([]record.AccountRecord, error)
	IsExistAccountEmail(ctx context.Context, email string) bool
	IsExistUsername(ctx context.Context, username string) bool
	VerifyCredential(ctx context.Context, username string) (record.AccountRecord, error)
	FindAccountUser(ctx context.Context, username string) (utils.UserAccounts, error)
	SaveLoginHistories(ctx context.Context, historiesRecord record.AccountLoginHistoriesRecord) error
	UpdateLogoutAt(ctx context.Context, userId int64, token string) error
}
