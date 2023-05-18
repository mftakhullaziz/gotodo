package tasks

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/tasks"
	"reflect"
	"testing"
)

func TestNewTaskRepositoryImpl(t *testing.T) {
	type args struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want tasks.TaskRecordRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskRepositoryImpl(tt.args.SQL, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepositoryImpl_DeleteTaskById(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx    context.Context
		taskId int64
		userId int64
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			if err := t.DeleteTaskById(tt.args.ctx, tt.args.taskId, tt.args.userId); (err != nil) != tt.wantErr {
				t1.Errorf("DeleteTaskById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskRepositoryImpl_FindTaskAll(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []record.TaskRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			got, err := t.FindTaskAll(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t1.Errorf("FindTaskAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("FindTaskAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepositoryImpl_FindTaskById(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx    context.Context
		taskId int64
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.TaskRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			got, err := t.FindTaskById(tt.args.ctx, tt.args.taskId, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t1.Errorf("FindTaskById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("FindTaskById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepositoryImpl_SaveTask(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx        context.Context
		taskRecord record.TaskRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.TaskRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			got, err := t.SaveTask(tt.args.ctx, tt.args.taskRecord)
			if (err != nil) != tt.wantErr {
				t1.Errorf("SaveTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SaveTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepositoryImpl_UpdateTask(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx        context.Context
		taskId     int64
		taskRecord record.TaskRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.TaskRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			got, err := t.UpdateTask(tt.args.ctx, tt.args.taskId, tt.args.taskRecord)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskRepositoryImpl_UpdateTaskStatus(t1 *testing.T) {
	type fields struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	type args struct {
		ctx       context.Context
		taskId    int64
		userId    int64
		completed string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.TaskRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskRepositoryImpl{
				SQL:      tt.fields.SQL,
				validate: tt.fields.validate,
			}
			got, err := t.UpdateTaskStatus(tt.args.ctx, tt.args.taskId, tt.args.userId, tt.args.completed)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateTaskStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdateTaskStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
