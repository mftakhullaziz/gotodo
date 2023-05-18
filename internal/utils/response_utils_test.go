package utils

import (
	"gotodo/internal/domain/models/response"
	"reflect"
	"testing"
)

func TestBuildAllResponseWithAuthorization(t *testing.T) {
	type args struct {
		handler   interface{}
		message   string
		totalData int
		requestAt string
	}
	tests := []struct {
		name string
		args args
		want response.DefaultServiceAllResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAllResponseWithAuthorization(tt.args.handler, tt.args.message, tt.args.totalData, tt.args.requestAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAllResponseWithAuthorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildEmptyResponse(t *testing.T) {
	type args struct {
		messages string
	}
	tests := []struct {
		name string
		args args
		want response.DefaultServiceResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildEmptyResponse(tt.args.messages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildEmptyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildResponseWithAuthorization(t *testing.T) {
	type args struct {
		handler    interface{}
		statusCode int
		taskId     int
		userId     string
		message1   string
	}
	tests := []struct {
		name string
		args args
		want response.DefaultServiceResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildResponseWithAuthorization(tt.args.handler, tt.args.statusCode, tt.args.taskId, tt.args.userId, tt.args.message1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildResponseWithAuthorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateResponses(t *testing.T) {
	type args struct {
		handler    interface{}
		statusCode int
		message1   string
		message2   string
	}
	tests := []struct {
		name string
		args args
		want response.DefaultServiceResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateResponses(tt.args.handler, tt.args.statusCode, tt.args.message1, tt.args.message2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateResponses() = %v, want %v", got, tt.want)
			}
		})
	}
}
