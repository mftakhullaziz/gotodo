package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
	errs "gotodo/internal/utils/errors"
)

type UserDetailUsecaseImpl struct {
	UserDetail accounts.UserService
	Validate   *validator.Validate
}

func NewUserDetailUsecaseImpl(userDetail accounts.UserService, validate *validator.Validate) account.UserDetailUsecase {
	return &UserDetailUsecaseImpl{UserDetail: userDetail, Validate: validate}
}

const formatDatetime = "2006-01-02 15:04:05"

func (u UserDetailUsecaseImpl) FindUserByUserIdUsecase(ctx context.Context, userId int64) (response.UserDetailResponse, error) {
	userDetailUsecase, err := u.UserDetail.FindUserByUserIdService(ctx, userId)
	errs.LoggerIfError(err)

	createTime := userDetailUsecase.CreatedAt.Format(formatDatetime)
	updateTime := userDetailUsecase.UpdatedAt.Format(formatDatetime)

	UserDetailResponse := response.UserDetailResponse{
		UserID:      userDetailUsecase.UserID,
		Username:    userDetailUsecase.Username,
		Password:    userDetailUsecase.Password,
		Email:       userDetailUsecase.Email,
		Name:        userDetailUsecase.Name,
		MobilePhone: userDetailUsecase.MobilePhone,
		Address:     userDetailUsecase.Address,
		Status:      userDetailUsecase.Status,
		CreatedAt:   createTime,
		UpdatedAt:   updateTime,
	}

	return UserDetailResponse, nil
}

func (u UserDetailUsecaseImpl) UpdateUserByUserIdUsecase(ctx context.Context, userId int64, request request.UserRequest) (response.UserDetailResponse, error) {
	err := u.Validate.Struct(request)
	errs.LoggerIfError(err)

	updateUser, errUsecase := u.UserDetail.UpdateUserByUserIdService(ctx, userId, request)
	errs.LoggerIfError(errUsecase)

	createTime := updateUser.CreatedAt.Format(formatDatetime)
	updateTime := updateUser.UpdatedAt.Format(formatDatetime)

	updateUserResult := response.UserDetailResponse{
		UserID:      updateUser.UserID,
		Username:    updateUser.Username,
		Password:    updateUser.Password,
		Email:       updateUser.Email,
		Name:        updateUser.Name,
		MobilePhone: updateUser.MobilePhone,
		Address:     updateUser.Address,
		Status:      updateUser.Status,
		CreatedAt:   createTime,
		UpdatedAt:   updateTime,
	}

	return updateUserResult, nil
}
