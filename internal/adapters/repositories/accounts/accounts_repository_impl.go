package accounts

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
	"gotodo/internal/utils"
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
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).Create(&accountRecord)
	if query.Error != nil {
		tx.Rollback()
		return record.AccountRecord{}, query.Error
	}
	tx.Commit()
	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountById(ctx context.Context, id int64) (record.AccountRecord, error) {
	var accountRecord record.AccountRecord
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).First(&accountRecord, id)
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.AccountRecord{}, ErrRecordNotFound
		}
		return record.AccountRecord{}, query.Error
	}
	tx.Commit()
	return accountRecord, nil
}

func (a AccountRepositoryImpl) UpdateAccount(ctx context.Context, id int64, accountRecord record.AccountRecord) (record.AccountRecord, error) {
	var existingAccount record.AccountRecord
	tx := a.SQL.Begin()
	// Check if the record exists
	query := tx.WithContext(ctx).First(&existingAccount, id).Error
	if query != nil {
		tx.Rollback()
		if errors.Is(query, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.AccountRecord{}, ErrRecordNotFound
		}
		return record.AccountRecord{}, query
	}

	// Update the fields of the existing record with the fields of the taskRecord argument
	qUpdate := tx.WithContext(ctx).Model(&existingAccount).Updates(accountRecord).Error
	if qUpdate != nil {
		tx.Rollback()
		return record.AccountRecord{}, qUpdate
	}

	tx.Commit()
	return existingAccount, nil
}

func (a AccountRepositoryImpl) DeleteAccountById(ctx context.Context, id int64) error {
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).Delete(&record.AccountRecord{}, id)
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return ErrRecordNotFound
		}
		return query.Error
	}
	tx.Commit()
	return nil
}

func (a AccountRepositoryImpl) FindAccountAll(ctx context.Context) ([]record.AccountRecord, error) {
	var accountRecords []record.AccountRecord
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).Find(&accountRecords)
	if query.Error != nil {
		tx.Rollback()
		return []record.AccountRecord{}, query.Error
	}
	tx.Commit()
	return accountRecords, nil
}

func (a AccountRepositoryImpl) IsExistAccountEmail(ctx context.Context, email string) bool {
	accountRecord := &record.UserDetailRecord{}
	tx := a.SQL.Begin()
	var count int64
	query := tx.WithContext(ctx).Model(accountRecord).Where("email = ?", email).Count(&count)
	if count != 0 {
		return true
	}
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	tx.Commit()
	return false
}

func (a AccountRepositoryImpl) IsExistUsername(ctx context.Context, username string) bool {
	accountRecord := &record.UserDetailRecord{}
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).
		Table("accounts").
		Joins("inner join user_details "+
			"on accounts.user_id = user_details.user_id "+
			"and accounts.username = user_details.username").
		Where("accounts.username = ?", username).
		Pluck("user_details.username", &accountRecord.Username)

	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	tx.Commit()

	return true
}

func (a AccountRepositoryImpl) VerifyCredential(ctx context.Context, username string) (record.AccountRecord, error) {
	accountRecord := record.AccountRecord{}
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).
		Select("username, password").
		Where("username = ? and status = ?", username, "active").
		First(&accountRecord)

	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return record.AccountRecord{}, errors.New("errors record not found")
		}
		return record.AccountRecord{}, query.Error
	} else if query.RowsAffected == 0 {
		tx.Rollback()
		return record.AccountRecord{}, nil
	}
	tx.Commit()

	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountUser(ctx context.Context, username string) (utils.UserAccounts, error) {
	userAccount := utils.UserAccounts{}
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).
		Joins("inner join user_details ud on accounts.user_id = ud.user_id and accounts.username = ud.username").
		Where("accounts.username = ? and accounts.status = ?", username, "active").
		First(&userAccount.Accounts)

	utils.StructJoinUserAccountRecordErrorUtils(query)

	queryFetchUser := tx.WithContext(ctx).Table("user_details").
		Where("user_id = ?", userAccount.Accounts.UserID).
		First(&userAccount.Users)

	utils.StructJoinUserAccountRecordErrorUtils(queryFetchUser)
	tx.Commit()

	return userAccount, nil
}

func (a AccountRepositoryImpl) SaveLoginHistories(ctx context.Context, historiesRecord record.AccountLoginHistoriesRecord) error {
	tx := a.SQL.Begin()
	query := tx.WithContext(ctx).Create(&historiesRecord)
	if query.Error != nil {
		tx.Rollback()
		return query.Error
	}
	tx.Commit()

	return nil
}

func (a AccountRepositoryImpl) UpdateLogoutAt(ctx context.Context, userId int64, token string) error {
	var historiesAccount record.AccountLoginHistoriesRecord

	tx := a.SQL.Begin()
	// Check if the record exists
	query := tx.WithContext(ctx).Where("user_id = ? AND token = ?", userId, token).First(&historiesAccount)
	if query != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors task record not found")
			return ErrRecordNotFound
		}
		return query.Error
	}

	// Save login histories
	historiesAccount.LoginOutAt = time.Now()
	querySave := tx.WithContext(ctx).Save(&historiesAccount)
	if querySave.Error != nil {
		tx.Rollback()
		return querySave.Error
	}
	tx.Commit()

	return nil
}
