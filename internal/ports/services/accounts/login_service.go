package accounts

import (
	"context"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
)

type LoginService interface {
	VerifyCredentialAccount(ctx context.Context, request request.LoginRequest) (dto.AccountDTO, string, error)
	LogoutAccountService(ctx context.Context, userId int64, token string) error
}
