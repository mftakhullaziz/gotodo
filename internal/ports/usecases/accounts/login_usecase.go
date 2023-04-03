package accounts

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type LoginUsecase interface {
	LoginAccountUseCase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error)
}
