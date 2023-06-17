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
	tx := t.SQL.Begin()
	query := tx.WithContext(ctx).Create(&taskRecord)
	if query.Error != nil {
		tx.Rollback()
		return record.TaskRecord{}, query.Error
	}
	tx.Commit()
	fmt.Println("result: ", taskRecord)
	return taskRecord, nil
}

func (t TaskRepositoryImpl) FindTaskById(ctx context.Context, taskId int64, userId int64) (record.TaskRecord, error) {
	var taskRecord record.TaskRecord
	tx := t.SQL.Begin()
	query := tx.WithContext(ctx).Where("user_id = ?", userId).First(&taskRecord, taskId)
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, query.Error
	}
	tx.Commit()
	return taskRecord, nil
}

func (t TaskRepositoryImpl) UpdateTask(ctx context.Context, taskId int64, taskRecord record.TaskRecord) (record.TaskRecord, error) {
	logger := utils.LoggerParent()
	var existingTask record.TaskRecord
	tx := t.SQL.Begin()
	// Check if the record exists
	query := tx.WithContext(ctx).First(&existingTask, taskId).Error
	if query != nil {
		tx.Rollback()
		if errors.Is(query, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, query
	}
	logger.Log.Infoln("Find Record Based On Task Id: ", existingTask)

	// Update the fields of the existing record with the fields of the taskRecord argument
	qUpdate := tx.WithContext(ctx).Model(&existingTask).Updates(taskRecord).Error
	if qUpdate != nil {
		tx.Rollback()
		return record.TaskRecord{}, qUpdate
	}
	logger.Log.Infoln("Find Record Based On Task Id After Update: ", existingTask)
	tx.Commit()

	return existingTask, nil
}

func (t TaskRepositoryImpl) DeleteTaskById(ctx context.Context, taskId int64, userId int64) error {
	var taskRecord record.TaskRecord
	tx := t.SQL.Begin()
	query := tx.WithContext(ctx).Where("user_id = ?", userId).First(&taskRecord, taskId)
	if query.Error != nil {
		tx.Rollback()
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors Record Not Found")
			return ErrRecordNotFound
		}
		return query.Error
	}

	// Set task status inactive
	taskRecord.TaskStatus = "inactive"
	// Update the fields of the existing record with the fields of the taskRecord argument
	qSave := tx.WithContext(ctx).Save(&taskRecord)
	if qSave != nil {
		tx.Rollback()
		return qSave.Error
	}
	tx.Commit()

	return nil
}

func (t TaskRepositoryImpl) FindTaskAll(ctx context.Context, userId int64) ([]record.TaskRecord, error) {
	var taskRecords []record.TaskRecord
	tx := t.SQL.Begin()
	query := tx.WithContext(ctx).Where("user_id = ?", userId).Find(&taskRecords)
	if query.Error != nil {
		tx.Rollback()
		return []record.TaskRecord{}, query.Error
	}
	tx.Commit()

	return taskRecords, nil
}

func (t TaskRepositoryImpl) UpdateTaskStatus(ctx context.Context, taskId int64, userId int64, completed string) (record.TaskRecord, error) {
	var taskRecord record.TaskRecord
	tx := t.SQL.Begin()
	// Check if the record exists
	query := tx.WithContext(ctx).Where("user_id = ?", userId).First(&taskRecord, taskId).Error
	if query != nil {
		tx.Rollback()
		if errors.Is(query, gorm.ErrRecordNotFound) {
			ErrRecordNotFound := errors.New("errors task record not found")
			return record.TaskRecord{}, ErrRecordNotFound
		}
		return record.TaskRecord{}, query
	}

	if completed == "" {
		log.Info("update task completed is failed")
		return record.TaskRecord{}, nil
	}

	// Set task status inactive
	taskRecord.Completed = true
	taskRecord.CompletedAt = time.Now()
	qUpdateCompleted := tx.WithContext(ctx).Save(&taskRecord)

	if qUpdateCompleted.Error != nil {
		tx.Rollback()
		return record.TaskRecord{}, qUpdateCompleted.Error
	}
	tx.Commit()

	return taskRecord, nil
}
