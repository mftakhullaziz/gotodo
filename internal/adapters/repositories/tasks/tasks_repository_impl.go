package tasks

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/tasks"
	"gotodo/internal/utils"
	"time"
)

type TaskRepositoryImpl struct {
	SQL      *gorm.DB
	validate *validator.Validate
}

func NewTaskRepositoryImpl(SQL *gorm.DB, validate *validator.Validate) tasks.TaskRecordRepository {
	return &TaskRepositoryImpl{SQL: SQL, validate: validate}
}

func (t TaskRepositoryImpl) SaveTask(ctx context.Context, taskRecord record.TaskRecord) (record.TaskRecord, error) {
	result := t.SQL.WithContext(ctx).Create(&taskRecord)
	if result.Error != nil {
		return record.TaskRecord{}, result.Error
	}
	fmt.Println("result: ", taskRecord)
	return taskRecord, nil
}

func (t TaskRepositoryImpl) FindTaskById(ctx context.Context, taskId int64, userId int64) (record.TaskRecord, error) {
	var taskRecord record.TaskRecord
	result := t.SQL.WithContext(ctx).Where("user_id = ?", userId).First(&taskRecord, taskId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, result.Error
	}
	return taskRecord, nil
}

func (t TaskRepositoryImpl) UpdateTask(ctx context.Context, taskId int64, taskRecord record.TaskRecord) (record.TaskRecord, error) {
	logger := utils.LoggerParent()
	var existingTask record.TaskRecord

	// Check if the record exists
	err := t.SQL.WithContext(ctx).First(&existingTask, taskId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, err
	}
	logger.Infoln("Find Record Based On Task Id: ", existingTask)

	// Update the fields of the existing record with the fields of the taskRecord argument
	err = t.SQL.WithContext(ctx).Model(&existingTask).Updates(taskRecord).Error
	if err != nil {
		return record.TaskRecord{}, err
	}
	logger.Infoln("Find Record Based On Task Id After Update: ", existingTask)

	return existingTask, nil
}

func (t TaskRepositoryImpl) DeleteTaskById(ctx context.Context, taskId int64, userId int64) error {
	var taskRecord record.TaskRecord
	result := t.SQL.WithContext(ctx).
		Where("user_id = ?", userId).
		First(&taskRecord, taskId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return ErrRecordNotFound
		}
		return result.Error
	}

	// Set task status inactive
	taskRecord.TaskStatus = "inactive"
	// Update the fields of the existing record with the fields of the taskRecord argument
	err := t.SQL.WithContext(ctx).Save(&taskRecord)
	if err != nil {
		return err.Error
	}

	return nil
}

func (t TaskRepositoryImpl) FindTaskAll(ctx context.Context, userId int64) ([]record.TaskRecord, error) {
	var taskRecords []record.TaskRecord
	result := t.SQL.WithContext(ctx).Where("user_id = ?", userId).Find(&taskRecords)
	if result.Error != nil {
		return []record.TaskRecord{}, result.Error
	}
	return taskRecords, nil
}

func (t TaskRepositoryImpl) UpdateTaskStatus(ctx context.Context, taskId int64, userId int64, completed string) (record.TaskRecord, error) {
	var taskRecord record.TaskRecord

	// Check if the record exists
	err := t.SQL.WithContext(ctx).Where("user_id = ?", userId).First(&taskRecord, taskId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors task record not found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, err
	}

	if completed == "" {
		log.Info("update task completed is failed")
		return record.TaskRecord{}, nil
	}

	// Set task status inactive
	taskRecord.Completed = true
	taskRecord.CompletedAt = time.Now()
	updateCompleted := t.SQL.WithContext(ctx).Save(&taskRecord)

	if updateCompleted.Error != nil {
		return record.TaskRecord{}, updateCompleted.Error
	}

	return taskRecord, nil
}
