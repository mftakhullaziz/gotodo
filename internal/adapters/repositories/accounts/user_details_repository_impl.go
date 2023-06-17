package accounts

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
)

type UserDetailRepositoryImpl struct {
	UserDetailRepository accounts.UserDetailRecordRepository
	SQL                  *gorm.DB
	validate             *validator.Validate
}

func NewUserDetailRepositoryImpl(SQL *gorm.DB, validate *validator.Validate) accounts.UserDetailRecordRepository {
	return &UserDetailRepositoryImpl{SQL: SQL, validate: validate}
}

func (u UserDetailRepositoryImpl) SaveUser(ctx context.Context, userRecord record.UserDetailRecord) (record.UserDetailRecord, error) {
	tx := u.SQL.Begin()
	query := tx.WithContext(ctx).Create(&userRecord)
	if query.Error != nil {
		tx.Rollback()
		return record.UserDetailRecord{}, query.Error
	}
	tx.Commit()
	return userRecord, nil
}

func (u UserDetailRepositoryImpl) FindUserById(ctx context.Context, userid int64) (record.UserDetailRecord, error) {
	var userRecord record.UserDetailRecord
	tx := u.SQL.Begin()
	query := tx.WithContext(ctx).First(&userRecord, userid)
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.UserDetailRecord{}, ErrRecordNotFound
		}
		return record.UserDetailRecord{}, query.Error
	}
	tx.Commit()

	return userRecord, nil
}

func (u UserDetailRepositoryImpl) UpdateUser(ctx context.Context, userId int64, userRecord record.UserDetailRecord) (record.UserDetailRecord, error) {
	var existingUser record.UserDetailRecord
	tx := u.SQL.Begin()

	// Check if the record exists
	query := tx.WithContext(ctx).First(&existingUser, userId).Error
	if query != nil {
		tx.Rollback()
		if errors.Is(query, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.UserDetailRecord{}, ErrRecordNotFound
		}
		return record.UserDetailRecord{}, query
	}

	// Update the fields of the existing record with the fields of the taskRecord argument
	qUpdate := tx.WithContext(ctx).Model(&existingUser).Updates(userRecord).Error
	if qUpdate != nil {
		tx.Rollback()
		return record.UserDetailRecord{}, qUpdate
	}
	tx.Commit()

	return existingUser, nil
}

func (u UserDetailRepositoryImpl) DeleteUserById(ctx context.Context, userId int64) error {
	tx := u.SQL.Begin()
	query := tx.WithContext(ctx).Delete(&record.UserDetailRecord{}, userId)
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

func (u UserDetailRepositoryImpl) FindUserAll(ctx context.Context) ([]record.UserDetailRecord, error) {
	var userRecords []record.UserDetailRecord
	tx := u.SQL.Begin()
	query := tx.WithContext(ctx).Find(&userRecords)
	if query.Error != nil {
		tx.Rollback()
		return []record.UserDetailRecord{}, query.Error
	}
	tx.Commit()

	return userRecords, nil
}
