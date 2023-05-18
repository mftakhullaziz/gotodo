package accounts

import (
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"net/http"
	"reflect"
	"testing"
)

func TestNewUserDetailHandlerAPI(t *testing.T) {
	type args struct {
		userUsecase accounts.UserDetailUsecase
	}
	tests := []struct {
		name string
		args args
		want api.UserHandlerAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDetailHandlerAPI(tt.args.userUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDetailHandlerAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailHandlerAPI_DeleteUserHandler(t *testing.T) {
	type fields struct {
		UserUsecase accounts.UserDetailUsecase
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
			u := UserDetailHandlerAPI{
				UserUsecase: tt.fields.UserUsecase,
			}
			u.DeleteUserHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestUserDetailHandlerAPI_FindDataUserDetailHandler(t *testing.T) {
	type fields struct {
		UserUsecase accounts.UserDetailUsecase
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
			u := UserDetailHandlerAPI{
				UserUsecase: tt.fields.UserUsecase,
			}
			u.FindDataUserDetailHandler(tt.args.writer, tt.args.requests)
		})
	}
}

func TestUserDetailHandlerAPI_UpdateUserDetailHandler(t *testing.T) {
	type fields struct {
		UserUsecase accounts.UserDetailUsecase
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
			u := UserDetailHandlerAPI{
				UserUsecase: tt.fields.UserUsecase,
			}
			u.UpdateUserDetailHandler(tt.args.writer, tt.args.requests)
		})
	}
}
