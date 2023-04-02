package accounts

import (
	"context"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
)

type RegisterUseCase interface {
	CreateAccountUseCase(ctx context.Context, request request.RegisterRequest) (response.RegisterResponse, error)
}
