package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
	"gotodo/internal/utils"
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
	logger := utils.LoggerParent()
	log := logger.Log

	validate := l.Validate.Struct(request)
	utils.PanicIfError(validate)

	credentialAccount, err := l.AccountRepository.VerifyCredential(ctx, request.Username)
	if err != nil {
		log.Info("username not found")
		return dto.AccountDTO{}, "", err
	}

	comparedPassword, errPassword := utils.ComparedPassword(credentialAccount.Password, []byte(request.Password))

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

	userRecord := findUserAccount.Users
	accountRecord := findUserAccount.Accounts

	tokenGenerate, expireTokenGenerate, errToken := middleware.GenerateJWTToken()
	utils.LoggerIfError(errToken)

	// Add token to cache
	err = middleware.GenerateTokenToCache(strconv.Itoa(int(findUserAccount.Users.UserID)), tokenGenerate, expireTokenGenerate)
	utils.LoggerIfError(err)

	// Add token authorization header
	err = middleware.MakeAuthenticatedRequest(tokenGenerate)
	utils.LoggerIfError(err)

	optionalLoginHistory := utils.NewOptionalColumnParams{
		BearerToken: tokenGenerate, TimeNow: time.Now(), TimeIsNull: time.Time{},
	}

	userLoginHistory := utils.UserAndAccountRecordToAccountLoginHistoryRecord(
		userRecord, accountRecord, optionalLoginHistory, expireTokenGenerate)

	saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
	utils.LoggerIfError(saveLoginHistory)

	userAccounts := utils.RecordToAccountDTO(accountRecord)

	return userAccounts, optionalLoginHistory.BearerToken, nil
}

func (l LoginServiceImpl) LogoutAccountService(ctx context.Context, userId int64, token string) error {
	err := utils.ValidateIntValue(int(userId))
	utils.LoggerIfError(err)

	logoutService := l.AccountRepository.UpdateLogoutAt(ctx, userId, token)
	utils.LoggerIfError(logoutService)

	return nil
}
