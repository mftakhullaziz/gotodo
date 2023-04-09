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
	result := u.SQL.WithContext(ctx).Create(&userRecord)
	if result.Error != nil {
		return record.UserDetailRecord{}, result.Error
	}
	return userRecord, nil
}

func (u UserDetailRepositoryImpl) FindUserById(ctx context.Context, userid int64) (record.UserDetailRecord, error) {
	var userRecord record.UserDetailRecord
	result := u.SQL.WithContext(ctx).First(&userRecord, userid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return record.UserDetailRecord{}, ErrRecordNotFound
		}
		return record.UserDetailRecord{}, result.Error
	}
	return userRecord, nil
}

func (u UserDetailRepositoryImpl) UpdateUser(ctx context.Context, userId int64, userRecord record.UserDetailRecord) (record.UserDetailRecord, error) {
	var existingUser record.UserDetailRecord

	// Check if the record exists
	err := u.SQL.WithContext(ctx).First(&existingUser, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return record.UserDetailRecord{}, ErrRecordNotFound
		}
		return record.UserDetailRecord{}, err
	}

	// Update the fields of the existing record with the fields of the taskRecord argument
	err = u.SQL.WithContext(ctx).Model(&existingUser).Updates(userRecord).Error
	if err != nil {
		return record.UserDetailRecord{}, err
	}

	return existingUser, nil
}

func (u UserDetailRepositoryImpl) DeleteUserById(ctx context.Context, userId int64) error {
	result := u.SQL.WithContext(ctx).Delete(&record.UserDetailRecord{}, userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("error Record Not Found")
			return ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}

func (u UserDetailRepositoryImpl) FindUserAll(ctx context.Context) ([]record.UserDetailRecord, error) {
	var userRecords []record.UserDetailRecord
	result := u.SQL.WithContext(ctx).Find(&userRecords)
	if result.Error != nil {
		return []record.UserDetailRecord{}, result.Error
	}
	return userRecords, nil
}
