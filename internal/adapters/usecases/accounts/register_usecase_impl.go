package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
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
	utils.PanicIfError(err)

	registerUseCase, errRegister := r.Register.CreateNewAccount(ctx, request)
	if errRegister != nil {
		return response.RegisterResponse{}, err
	}

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
