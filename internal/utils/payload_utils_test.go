package utils

import (
	"net/http"
	"testing"
)

func TestReadFromRequestBody(t *testing.T) {
	type args struct {
		request *http.Request
		result  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFromRequestBody(tt.args.request, tt.args.result)
		})
	}
}

func TestWriteToResponseBody(t *testing.T) {
	type args struct {
		writer   http.ResponseWriter
		response interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteToResponseBody(tt.args.writer, tt.args.response)
		})
	}
}
