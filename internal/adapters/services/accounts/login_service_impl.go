package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
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
	log := helpers.LoggerParent()

	validate := l.Validate.Struct(request)
	helpers.PanicIfError(validate)

	credentialAccount, err := l.AccountRepository.VerifyCredential(ctx, request.Username)
	if err != nil {
		log.Info("username not found")
		return dto.AccountDTO{}, "", err
	}
	log.Info("credential: ", credentialAccount.Username, ", ", credentialAccount.Password)

	comparedPassword, errPassword := helpers.
		ComparedPassword(credentialAccount.Password, []byte(request.Password))

	// validate password is matched
	if comparedPassword != true && errPassword != nil {
		log.Info("password not matched")
		return dto.AccountDTO{}, "", errPassword
	}
	// if validate password and username
	findUserAccount, errUserAccount := l.AccountRepository.FindAccountUser(ctx, credentialAccount.Username)
	if errUserAccount != nil {
		log.Error("error find user account not found", errUserAccount.Error())
	}
	log.Info("user account find: ", findUserAccount)

	userRecord := findUserAccount.Users
	accountRecord := findUserAccount.Accounts
	log.Infoln("account record: ", accountRecord)

	tokenGenerate, expireTokenGenerate, errToken := middleware.GenerateJWTToken()
	helpers.LoggerIfError(errToken)

	// Add token to cache
	err = middleware.GenerateTokenToCache(strconv.Itoa(int(findUserAccount.Users.UserID)), tokenGenerate, expireTokenGenerate)
	helpers.LoggerIfError(err)

	// Add token authorization header
	err = middleware.MakeAuthenticatedRequest(tokenGenerate)
	helpers.LoggerIfError(err)

	optionalLoginHistory := helpers.NewOptionalColumnParams{
		BearerToken: tokenGenerate, TimeNow: time.Now(), TimeIsNull: time.Time{},
	}

	userLoginHistory := helpers.UserAndAccountRecordToAccountLoginHistoryRecord(
		userRecord, accountRecord, optionalLoginHistory, expireTokenGenerate)

	saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
	helpers.LoggerIfError(saveLoginHistory)

	userAccounts := helpers.RecordToAccountDTO(accountRecord)

	return userAccounts, optionalLoginHistory.BearerToken, nil
}

func (l LoginServiceImpl) LogoutAccountService(ctx context.Context, userId int64, token string) error {
	err := helpers.ValidateIntValue(int(userId))
	helpers.LoggerIfError(err)

	logoutService := l.AccountRepository.UpdateLogoutAt(ctx, userId, token)
	helpers.LoggerIfError(logoutService)

	return nil
}
