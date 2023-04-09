package accounts

import (
	"context"
	"gotodo/internal/domain/dto"
)

type UserService interface {
	FindUserByUserIdService(ctx context.Context, userId int64) (dto.UserDetailDTO, error)
	UpdateUserByUserIdService(ctx context.Context, userId int64) (dto.UserDetailDTO, error)
}
