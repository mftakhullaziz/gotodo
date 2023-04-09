package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
)

type UserDetailUsecaseImpl struct {
	UserDetail accounts.UserService
	Validate   *validator.Validate
}

func NewUserDetailUsecaseImpl(userDetail accounts.UserService, validate *validator.Validate) account.UserDetailUsecase {
	return &UserDetailUsecaseImpl{UserDetail: userDetail, Validate: validate}
}

func (u UserDetailUsecaseImpl) FindUserByUserIdUsecase(ctx context.Context, userId int64) (response.UserDetailResponse, error) {
	UserDetailUsecase, err := u.UserDetail.FindUserByUserIdService(ctx, userId)
	helpers.LoggerIfError(err)

	UserDetailResponse := response.UserDetailResponse{
		UserID:      UserDetailUsecase.UserID,
		Username:    UserDetailUsecase.Username,
		Password:    UserDetailUsecase.Password,
		Email:       UserDetailUsecase.Email,
		Name:        UserDetailUsecase.Name,
		MobilePhone: UserDetailUsecase.MobilePhone,
		Address:     UserDetailUsecase.Address,
		Status:      UserDetailUsecase.Status,
		CreatedAt:   UserDetailUsecase.CreatedAt,
		UpdatedAt:   UserDetailUsecase.UpdatedAt}

	return UserDetailResponse, nil
}

func (u UserDetailUsecaseImpl) UpdateUserByUserIdUsecase(ctx context.Context, userId int64) (response.UserDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}
