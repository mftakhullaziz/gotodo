package accounts

import (
	"context"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
)

type RegisterService interface {
	CreateNewAccount(ctx context.Context, request request.RegisterRequest) (dto.CreateAccountDTO, error)
}
