package accounts

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
)

type AccountRepositoryImpl struct {
	AccountsRepository accounts.AccountRecordRepository
	SQL                *gorm.DB
	validate           *validator.Validate
}

func NewAccountsRepositoryImpl(SQL *gorm.DB, validate *validator.Validate) accounts.AccountRecordRepository {
	return &AccountRepositoryImpl{SQL: SQL, validate: validate}
}

func (a AccountRepositoryImpl) SaveAccount(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error) {
	result := a.SQL.WithContext(ctx).Create(&accountRecord)
	if result.Error != nil {
		return record.AccountRecord{}, result.Error
	}
	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountById(ctx context.Context, id int64) (record.AccountRecord, error) {
	var accountRecord record.AccountRecord
	result := a.SQL.WithContext(ctx).First(&accountRecord, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return record.AccountRecord{}, ErrRecordNotFound
		}
		return record.AccountRecord{}, result.Error
	}
	return accountRecord, nil
}

func (a AccountRepositoryImpl) UpdateAccount(ctx context.Context, id int64, accountRecord record.AccountRecord) (record.AccountRecord, error) {
	var existingAccount record.AccountRecord

	// Check if the record exists
	err := a.SQL.WithContext(ctx).First(&existingAccount, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return record.AccountRecord{}, ErrRecordNotFound
		}
		return record.AccountRecord{}, err
	}

	// Update the fields of the existing record with the fields of the taskRecord argument
	err = a.SQL.WithContext(ctx).Model(&existingAccount).Updates(accountRecord).Error
	if err != nil {
		return record.AccountRecord{}, err
	}

	return existingAccount, nil
}

func (a AccountRepositoryImpl) DeleteAccountById(ctx context.Context, id int64) error {
	result := a.SQL.WithContext(ctx).Delete(&record.AccountRecord{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}

func (a AccountRepositoryImpl) FindAccountAll(ctx context.Context) ([]record.AccountRecord, error) {
	var accountRecords []record.AccountRecord
	result := a.SQL.WithContext(ctx).Find(&accountRecords)
	if result.Error != nil {
		return []record.AccountRecord{}, result.Error
	}
	return accountRecords, nil
}

func (a AccountRepositoryImpl) FindAccountByEmail(ctx context.Context, email string) bool {
	accountRecord := &record.UserDetailRecord{}
	result := a.SQL.WithContext(ctx).Where("email = ?", email).First(accountRecord)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

func (a AccountRepositoryImpl) IsDuplicateUsername(ctx context.Context, username string) bool {
	accountRecord := &record.UserDetailRecord{}
	result := a.SQL.WithContext(ctx).
		Table("accounts").
		Joins("inner join user_details "+
			"on accounts.user_id = user_details.user_id "+
			"and accounts.username = user_details.username").
		Where("accounts.username = ?", username).
		Pluck("user_details.username", &accountRecord.Username)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
