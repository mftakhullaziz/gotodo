package accounts

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type LoginUsecase interface {
	LoginAccountUsecase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error)
	LogoutAccountUsecase(ctx context.Context, userId int, token string) error
}
