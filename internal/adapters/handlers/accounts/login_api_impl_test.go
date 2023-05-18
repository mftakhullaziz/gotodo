package accounts

import (
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"net/http"
	"reflect"
	"testing"
)

func TestLoginHandlerAPI_LoginHandler(t *testing.T) {
	type fields struct {
		LoginUsecase accounts.LoginUsecase
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
		t.Run(tt.name, func(t *testing.T) {
			l := LoginHandlerAPI{
				LoginUsecase: tt.fields.LoginUsecase,
			}
			l.LoginHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestLoginHandlerAPI_LogoutHandler(t *testing.T) {
	type fields struct {
		LoginUsecase accounts.LoginUsecase
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
		t.Run(tt.name, func(t *testing.T) {
			l := LoginHandlerAPI{
				LoginUsecase: tt.fields.LoginUsecase,
			}
			l.LogoutHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestNewLoginHandlerAPI(t *testing.T) {
	type args struct {
		loginUsecase accounts.LoginUsecase
	}
	tests := []struct {
		name string
		args args
		want api.LoginHandlerAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginHandlerAPI(tt.args.loginUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginHandlerAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
