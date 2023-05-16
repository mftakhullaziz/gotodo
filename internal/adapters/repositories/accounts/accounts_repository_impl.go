package accounts

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/struct"
	"time"
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
			ErrRecordNotFound := errors.New("errors Record Not Found")
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
			ErrRecordNotFound := errors.New("errors Record Not Found")
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
			ErrRecordNotFound := errors.New("errors Record Not Found")
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

func (a AccountRepositoryImpl) IsExistAccountEmail(ctx context.Context, email string) bool {
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

func (a AccountRepositoryImpl) IsExistUsername(ctx context.Context, username string) bool {
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

func (a AccountRepositoryImpl) VerifyCredential(ctx context.Context, username string) (record.AccountRecord, error) {
	accountRecord := record.AccountRecord{}
	result := a.SQL.WithContext(ctx).Select("username, password").
		Where("username = ? and status = ?", username, "active").
		First(&accountRecord)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return record.AccountRecord{}, errors.New("errors record not found")
		}
		return record.AccountRecord{}, result.Error
	} else if result.RowsAffected == 0 {
		return record.AccountRecord{}, nil
	}

	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountUser(ctx context.Context, username string) (_struct.UserAccounts, error) {
	userAccount := _struct.UserAccounts{}

	resultAccount := a.SQL.WithContext(ctx).
		Joins("inner join user_details ud on accounts.user_id = ud.user_id and accounts.username = ud.username").
		Where("accounts.username = ? and accounts.status = ?", username, "active").
		First(&userAccount.Accounts)

	errs.StructJoinUserAccountRecordErrorUtils(resultAccount)

	resultUser := a.SQL.WithContext(ctx).Table("user_details").
		Where("user_id = ?", userAccount.Accounts.UserID).
		First(&userAccount.Users)

	errs.StructJoinUserAccountRecordErrorUtils(resultUser)

	return userAccount, nil
}

func (a AccountRepositoryImpl) SaveLoginHistories(ctx context.Context, historiesRecord record.AccountLoginHistoriesRecord) error {
	result := a.SQL.WithContext(ctx).Create(&historiesRecord)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a AccountRepositoryImpl) UpdateLogoutAt(ctx context.Context, userId int64, token string) error {
	var historiesAccount record.AccountLoginHistoriesRecord

	// Check if the record exists
	err := a.SQL.WithContext(ctx).Where("user_id = ? AND token = ?", userId, token).First(&historiesAccount).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors task record not found")
			return ErrRecordNotFound
		}
		return err
	}

	historiesAccount.LoginOutAt = time.Now()
	saveHistoriesLogin := a.SQL.WithContext(ctx).Save(&historiesAccount)
	if saveHistoriesLogin.Error != nil {
		return saveHistoriesLogin.Error
	}

	return nil
}
