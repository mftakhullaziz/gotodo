package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
)

type RegisterUseCaseImpl struct {
	Register accounts.RegisterService
	Validate *validator.Validate
}

func NewRegisterUseCaseImpl(register accounts.RegisterService, validate *validator.Validate) account.RegisterUseCase {
	return &RegisterUseCaseImpl{Register: register, Validate: validate}
}

func (r RegisterUseCaseImpl) CreateAccountUseCase(ctx context.Context, request request.RegisterRequest) (response.RegisterResponse, error) {
	err := r.Validate.Struct(request)
	helpers.PanicIfError(err)

	registerUseCase, err := r.Register.CreateNewAccount(ctx, request)
	helpers.PanicIfError(err)

	result := response.RegisterResponse{
		AccountID: registerUseCase.AccountID,
		UserID:    registerUseCase.UserID,
		Username:  registerUseCase.Username,
		Password:  registerUseCase.Password,
		Status:    registerUseCase.Status,
		CreatedAt: registerUseCase.CreatedAt,
		UpdatedAt: registerUseCase.UpdatedAt,
	}

	return result, nil
}
