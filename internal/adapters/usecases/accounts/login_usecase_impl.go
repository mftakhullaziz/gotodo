package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
	"time"
)

type LoginUsecaseImpl struct {
	LoginService accounts.LoginService
	Validate     *validator.Validate
}

func NewLoginUsecaseImpl(loginService accounts.LoginService, validate *validator.Validate) account.LoginUsecase {
	return &LoginUsecaseImpl{LoginService: loginService, Validate: validate}
}

func (l LoginUsecaseImpl) LoginAccountUseCase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error) {
	err := l.Validate.Struct(request)
	helpers.PanicIfError(err)

	loginUsecase, errLogin := l.LoginService.VerifyCredentialAccount(ctx, request)
	if errLogin != nil {
		return response.LoginResponse{}, err
	}

	responseUsecase := response.LoginResponse{
		AccountID: loginUsecase.AccountID,
		Username:  loginUsecase.Username,
		Password:  loginUsecase.Password,
		LoginAt:   time.Time{},
	}
	return responseUsecase, nil
}
