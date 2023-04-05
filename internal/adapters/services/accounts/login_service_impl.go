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
	helpers.PanicIfErrorWithCustomMessage(err, "username not found")
	log.Info("verify credential: ", credentialAccount)

	comparedPassword, errCompare := helpers.ComparedPassword(credentialAccount.Password, []byte(request.Password))
	if comparedPassword != true && errCompare != nil {
		log.Info("password not matched")
		return dto.AccountDTO{}, "", errCompare
	} else {
		log.Info("user and password is validate")

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
		errCache := middleware.GenerateTokenToCache(strconv.Itoa(int(findUserAccount.Users.UserID)),
			tokenGenerate, expireTokenGenerate)
		helpers.LoggerIfError(errCache)

		// Add token authorization header
		header, errHeader := middleware.MakeAuthenticatedRequest(tokenGenerate)
		helpers.LoggerIfError(errHeader)
		log.Info("add to header authorization: ", header)

		optionalUserLoginHistory := helpers.NewOptionalColumnParams{
			Token:   tokenGenerate,
			TimeAt:  time.Now(),
			TimeOut: time.Time{},
		}

		userLoginHistory := helpers.UserAndAccountRecordToAccountLoginHistoryRecord(
			userRecord, accountRecord, optionalUserLoginHistory)
		log.Info("user login histories: ", userLoginHistory)

		saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
		helpers.LoggerIfError(saveLoginHistory)

		userAccounts := helpers.RecordToAccountDTO(accountRecord)
		log.Info("user accounts record to dto: ", userAccounts)

		return userAccounts, optionalUserLoginHistory.Token, nil
	}
}
