package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
	"gotodo/internal/utils"
	"reflect"
	"testing"
)

func TestAccountRepositoryImpl_DeleteAccountById(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			if err := a.DeleteAccountById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccountById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepositoryImpl_FindAccountAll(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []record.AccountRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.FindAccountAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAccountAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAccountAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_FindAccountById(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.AccountRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.FindAccountById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAccountById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAccountById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_FindAccountUser(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    utils.UserAccounts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.FindAccountUser(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAccountUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAccountUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_IsExistAccountEmail(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			if got := a.IsExistAccountEmail(tt.args.ctx, tt.args.email); got != tt.want {
				t.Errorf("IsExistAccountEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_IsExistUsername(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			if got := a.IsExistUsername(tt.args.ctx, tt.args.username); got != tt.want {
				t.Errorf("IsExistUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_SaveAccount(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx           context.Context
		accountRecord record.AccountRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.AccountRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.SaveAccount(tt.args.ctx, tt.args.accountRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_SaveLoginHistories(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx             context.Context
		historiesRecord record.AccountLoginHistoriesRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			if err := a.SaveLoginHistories(tt.args.ctx, tt.args.historiesRecord); (err != nil) != tt.wantErr {
				t.Errorf("SaveLoginHistories() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepositoryImpl_UpdateAccount(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx           context.Context
		id            int64
		accountRecord record.AccountRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.AccountRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.UpdateAccount(tt.args.ctx, tt.args.id, tt.args.accountRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepositoryImpl_UpdateLogoutAt(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx    context.Context
		userId int64
		token  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			if err := a.UpdateLogoutAt(tt.args.ctx, tt.args.userId, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLogoutAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountRepositoryImpl_VerifyCredential(t *testing.T) {
	type fields struct {
		AccountsRepository accounts.AccountRecordRepository
		SQL                *gorm.DB
		validate           *validator.Validate
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.AccountRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AccountRepositoryImpl{
				AccountsRepository: tt.fields.AccountsRepository,
				SQL:                tt.fields.SQL,
				validate:           tt.fields.validate,
			}
			got, err := a.VerifyCredential(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyCredential() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAccountsRepositoryImpl(t *testing.T) {
	type args struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want accounts.AccountRecordRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountsRepositoryImpl(tt.args.SQL, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountsRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
