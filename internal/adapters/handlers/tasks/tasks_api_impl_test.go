package tasks

import (
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/tasks"
	"net/http"
	"reflect"
	"testing"
)

func TestNewTaskHandlerAPI(t *testing.T) {
	type args struct {
		taskUseCase tasks.TaskUseCase
	}
	tests := []struct {
		name string
		args args
		want api.TaskHandlerAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskHandlerAPI(tt.args.taskUseCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskHandlerAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandlerAPI_CreateTaskHandler(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.CreateTaskHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestTaskHandlerAPI_DeleteTaskHandler(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.DeleteTaskHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestTaskHandlerAPI_FindTaskHandler(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.FindTaskHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestTaskHandlerAPI_FindTaskHandlerById(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.FindTaskHandlerById(tt.args.writer, tt.args.requests)
		})
	}
}

func TestTaskHandlerAPI_UpdateTaskHandler(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.UpdateTaskHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestTaskHandlerAPI_UpdateTaskStatusHandler(t1 *testing.T) {
	type fields struct {
		TaskUseCase tasks.TaskUseCase
	}
	type args struct {
		writer   http.ResponseWriter
		requests *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TaskHandlerAPI{
				TaskUseCase: tt.fields.TaskUseCase,
			}
			t.UpdateTaskStatusHandler(tt.args.writer, tt.args.requests)
		})
	}
}
