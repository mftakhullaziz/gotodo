package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/domain/models/response"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/services/accounts"
	account "gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
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
	utils.PanicIfError(err)

	loginUsecase, token, errLogin := l.LoginService.VerifyCredentialAccount(ctx, request)
	if errLogin != nil {
		return response.LoginResponse{}, err
	}

	responseUsecase := response.LoginResponse{
		AccountID:         int(loginUsecase.AccountID),
		Username:          loginUsecase.Username,
		Password:          loginUsecase.Password,
		LoginCreationTime: loginUsecase.CreatedAt,
		LoginToken:        token,
	}
	return responseUsecase, nil
}

func (l LoginUsecaseImpl) LogoutAccountUsecase(ctx context.Context, userId int, token string) error {
	err := utils.ValidateIntValue(userId)
	utils.LoggerIfError(err)

	// Check token in jwt or caching or redis if any remove from that
	// Logout first rules remove token from jwt
	tokenJwt := middleware.CheckAndRemoveExpiredToken(token)
	if tokenJwt == false {
		// Remove token from cache
		err = middleware.CheckAndRemoveTokenFromCache(token)
		if err != nil {
			panic(err.Error())
		}
		// Update to table histories logout
		logoutUsecase := l.LoginService.LogoutAccountService(ctx, int64(userId), token)
		utils.LoggerIfError(logoutUsecase)
	}

	return nil
}
