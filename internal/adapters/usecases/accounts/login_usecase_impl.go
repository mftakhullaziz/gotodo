package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
	errs "gotodo/internal/utils/errors"
	validate "gotodo/internal/utils/validator"
	"time"
)

type LoginUsecaseImpl struct {
	LoginService accounts.LoginService
	Validate     *validator.Validate
}

func NewLoginUsecaseImpl(loginService accounts.LoginService, validate *validator.Validate) account.LoginUsecase {
	return &LoginUsecaseImpl{LoginService: loginService, Validate: validate}
}

func (l LoginUsecaseImpl) LoginAccountUsecase(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error) {
	err := l.Validate.Struct(request)
	errs.PanicIfError(err)

	loginUsecase, token, errLogin := l.LoginService.VerifyCredentialAccount(ctx, request)
	if errLogin != nil {
		return response.LoginResponse{}, err
	}

	responseUsecase := response.LoginResponse{
		AccountID: loginUsecase.AccountID,
		Username:  loginUsecase.Username,
		Password:  loginUsecase.Password,
		LoginAt:   time.Time{},
		Token:     token,
	}
	return responseUsecase, nil
}

func (l LoginUsecaseImpl) LogoutAccountUsecase(ctx context.Context, userId int, token string) error {
	err := validate.ValidateIntValue(userId)
	errs.LoggerIfError(err)

	logoutUsecase := l.LoginService.LogoutAccountService(ctx, int64(userId), token)
	errs.LoggerIfError(logoutUsecase)

	return nil
}
