package register

//import (
//	"gotodo/internal/ports/handlers/api"
//	"gotodo/internal/ports/usecases/accounts"
//	"net/http"
//	"reflect"
//	"testing"
//)
//
//func TestNewRegisterHandlerAPI(t *testing.T) {
//	type args struct {
//		registerUseCase accounts.RegisterUseCase
//	}
//	tests := []struct {
//		name string
//		args args
//		want api.RegisterHandlers
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewRegisterHandlerAPI(tt.args.registerUseCase); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewRegisterHandlerAPI() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestRegisterHandlerAPI_ForgotPasswordHandler(t *testing.T) {
//	type fields struct {
//		RegisterUseCase accounts.RegisterUseCase
//	}
//	type args struct {
//		writer   http.ResponseWriter
//		requests *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := RegisterHandlers{
//				RegisterUseCase: tt.fields.RegisterUseCase,
//			}
//			r.ForgotPasswordHandler(tt.args.writer, tt.args.requests)
//		})
//	}
//}
//
//func TestRegisterHandlerAPI_RegisterHandler(t *testing.T) {
//	type fields struct {
//		RegisterUseCase accounts.RegisterUseCase
//	}
//	type args struct {
//		writer   http.ResponseWriter
//		requests *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := RegisterHandlers{
//				RegisterUseCase: tt.fields.RegisterUseCase,
//			}
//			r.RegisterHandler(tt.args.writer, tt.args.requests)
//		})
//	}
//}
