package accounts

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type UserDetailUsecase interface {
	FindUserByUserIdUsecase(ctx context.Context, userId int64) (response.UserDetailResponse, error)
	UpdateUserByUserIdUsecase(ctx context.Context, userId int64, request request.UserRequest) (response.UserDetailResponse, error)
}
