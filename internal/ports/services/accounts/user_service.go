package accounts

import (
	"context"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
)

type UserService interface {
	FindUserByUserIdService(ctx context.Context, userId int64) (dto.UserDetailDTO, error)
	UpdateUserByUserIdService(ctx context.Context, userId int64, request request.UserRequest) (dto.UserDetailDTO, error)
}
