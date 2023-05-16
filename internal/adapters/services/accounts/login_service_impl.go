package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
	"gotodo/internal/utils/converter"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"gotodo/internal/utils/password"
	validate "gotodo/internal/utils/validator"
	"strconv"
	"time"
)

type LoginServiceImpl struct {
	AccountRepository accounts.AccountRecordRepository
	Validate          *validator.Validate
}

func NewLoginServiceImpl(accountRepository accounts.AccountRecordRepository, validate *validator.Validate) account.LoginService {
	return &LoginServiceImpl{AccountRepository: accountRepository, Validate: validate}
}

func (l LoginServiceImpl) VerifyCredentialAccount(ctx context.Context, request request.LoginRequest) (dto.AccountDTO, string, error) {
	log := logger.LoggerParent()

	validate := l.Validate.Struct(request)
	errs.PanicIfError(validate)

	credentialAccount, err := l.AccountRepository.VerifyCredential(ctx, request.Username)
	if err != nil {
		log.Info("username not found")
		return dto.AccountDTO{}, "", err
	}
	log.Info("credential: ", credentialAccount.Username, ", ", credentialAccount.Password)

	comparedPassword, errPassword := password.ComparedPassword(credentialAccount.Password, []byte(request.Password))

	// validate password is matched
	if comparedPassword != true && errPassword != nil {
		log.Info("password not matched")
		return dto.AccountDTO{}, "", errPassword
	}
	// if validate password and username
	findUserAccount, errUserAccount := l.AccountRepository.FindAccountUser(ctx, credentialAccount.Username)
	if errUserAccount != nil {
		log.Error("errors find user account not found", errUserAccount.Error())
	}
	log.Info("user account find: ", findUserAccount)

	userRecord := findUserAccount.Users
	accountRecord := findUserAccount.Accounts
	log.Infoln("account record: ", accountRecord)

	tokenGenerate, expireTokenGenerate, errToken := middleware.GenerateJWTToken()
	errs.LoggerIfError(errToken)

	// Add token to cache
	err = middleware.GenerateTokenToCache(strconv.Itoa(int(findUserAccount.Users.UserID)), tokenGenerate, expireTokenGenerate)
	errs.LoggerIfError(err)

	// Add token authorization header
	err = middleware.MakeAuthenticatedRequest(tokenGenerate)
	errs.LoggerIfError(err)

	optionalLoginHistory := converter.NewOptionalColumnParams{
		BearerToken: tokenGenerate, TimeNow: time.Now(), TimeIsNull: time.Time{},
	}

	userLoginHistory := converter.UserAndAccountRecordToAccountLoginHistoryRecord(
		userRecord, accountRecord, optionalLoginHistory, expireTokenGenerate)

	saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
	errs.LoggerIfError(saveLoginHistory)

	userAccounts := converter.RecordToAccountDTO(accountRecord)

	return userAccounts, optionalLoginHistory.BearerToken, nil
}

func (l LoginServiceImpl) LogoutAccountService(ctx context.Context, userId int64, token string) error {
	err := validate.ValidateIntValue(int(userId))
	errs.LoggerIfError(err)

	logoutService := l.AccountRepository.UpdateLogoutAt(ctx, userId, token)
	errs.LoggerIfError(logoutService)

	return nil
}
