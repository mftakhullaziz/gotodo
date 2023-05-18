package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gotodo/internal/persistence/record"
	"gotodo/internal/ports/repositories/accounts"
	"reflect"
	"testing"
)

func TestNewUserDetailRepositoryImpl(t *testing.T) {
	type args struct {
		SQL      *gorm.DB
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want accounts.UserDetailRecordRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDetailRepositoryImpl(tt.args.SQL, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDetailRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailRepositoryImpl_DeleteUserById(t *testing.T) {
	type fields struct {
		UserDetailRepository accounts.UserDetailRecordRepository
		SQL                  *gorm.DB
		validate             *validator.Validate
	}
	type args struct {
		ctx    context.Context
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
		t.Run(tt.name, func(t *testing.T) {
			u := UserDetailRepositoryImpl{
				UserDetailRepository: tt.fields.UserDetailRepository,
				SQL:                  tt.fields.SQL,
				validate:             tt.fields.validate,
			}
			if err := u.DeleteUserById(tt.args.ctx, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDetailRepositoryImpl_FindUserAll(t *testing.T) {
	type fields struct {
		UserDetailRepository accounts.UserDetailRecordRepository
		SQL                  *gorm.DB
		validate             *validator.Validate
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []record.UserDetailRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserDetailRepositoryImpl{
				UserDetailRepository: tt.fields.UserDetailRepository,
				SQL:                  tt.fields.SQL,
				validate:             tt.fields.validate,
			}
			got, err := u.FindUserAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailRepositoryImpl_FindUserById(t *testing.T) {
	type fields struct {
		UserDetailRepository accounts.UserDetailRecordRepository
		SQL                  *gorm.DB
		validate             *validator.Validate
	}
	type args struct {
		ctx    context.Context
		userid int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.UserDetailRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserDetailRepositoryImpl{
				UserDetailRepository: tt.fields.UserDetailRepository,
				SQL:                  tt.fields.SQL,
				validate:             tt.fields.validate,
			}
			got, err := u.FindUserById(tt.args.ctx, tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailRepositoryImpl_SaveUser(t *testing.T) {
	type fields struct {
		UserDetailRepository accounts.UserDetailRecordRepository
		SQL                  *gorm.DB
		validate             *validator.Validate
	}
	type args struct {
		ctx        context.Context
		userRecord record.UserDetailRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.UserDetailRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserDetailRepositoryImpl{
				UserDetailRepository: tt.fields.UserDetailRepository,
				SQL:                  tt.fields.SQL,
				validate:             tt.fields.validate,
			}
			got, err := u.SaveUser(tt.args.ctx, tt.args.userRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDetailRepositoryImpl_UpdateUser(t *testing.T) {
	type fields struct {
		UserDetailRepository accounts.UserDetailRecordRepository
		SQL                  *gorm.DB
		validate             *validator.Validate
	}
	type args struct {
		ctx        context.Context
		userId     int64
		userRecord record.UserDetailRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    record.UserDetailRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserDetailRepositoryImpl{
				UserDetailRepository: tt.fields.UserDetailRepository,
				SQL:                  tt.fields.SQL,
				validate:             tt.fields.validate,
			}
			got, err := u.UpdateUser(tt.args.ctx, tt.args.userId, tt.args.userRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
