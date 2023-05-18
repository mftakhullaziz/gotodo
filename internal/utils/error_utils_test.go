package utils

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestFatalIfErrorWithCustomMessage(t *testing.T) {
	type args struct {
		err error
		log *logrus.Logger
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FatalIfErrorWithCustomMessage(tt.args.err, tt.args.log, tt.args.str)
		})
	}
}

func TestInternalServerError(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		err interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InternalServerError(tt.args.w, tt.args.r, tt.args.err)
		})
	}
}

func TestLoggerIfError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerIfError(tt.args.err)
		})
	}
}

func TestLoggerIfErrorWithCustomMessage(t *testing.T) {
	type args struct {
		err error
		log *logrus.Logger
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerIfErrorWithCustomMessage(tt.args.err, tt.args.log, tt.args.str)
		})
	}
}

func TestPanicIfError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PanicIfError(tt.args.err)
		})
	}
}

func TestPanicIfErrorWithCustomMessage(t *testing.T) {
	type args struct {
		err error
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PanicIfErrorWithCustomMessage(tt.args.err, tt.args.str)
		})
	}
}

func TestStructJoinUserAccountRecordErrorUtils(t *testing.T) {
	type args struct {
		gdb *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructJoinUserAccountRecordErrorUtils(tt.args.gdb)
		})
	}
}
