package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
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
	// log := helpers.LoggerParent()

	validate := l.Validate.Struct(request)
	helpers.PanicIfError(validate)

	credentialAccount, err := l.AccountRepository.VerifyCredential(ctx, request.Username)
	helpers.PanicIfErrorWithCustomMessage(err, "Username not found")

	comparedPassword, errCompare := helpers.ComparedPassword(credentialAccount.Password, []byte(request.Password))
	if comparedPassword != true && errCompare != nil {
		log.Info("Password not matched")
		return dto.AccountDTO{}, "", errCompare
	} else {
		log.Info("Password is matched")

		findUserAccount, errUserAccount := l.AccountRepository.FindAccountUser(ctx, credentialAccount.Username)
		if errUserAccount != nil {
			log.Error("Error find user account not found", errUserAccount.Error())
		}
		log.Info("User Account Find: ", findUserAccount)

		userRecord := findUserAccount.Users
		accountRecord := findUserAccount.Accounts

		tokenGenerate, expireTokenGenerate, errToken := middleware.GenerateJWTToken()
		helpers.LoggerIfError(errToken)

		// Add token to cache
		errCache := middleware.GenerateTokenToCache(strconv.Itoa(int(findUserAccount.Users.UserID)),
			tokenGenerate, expireTokenGenerate)
		helpers.LoggerIfError(errCache)

		// Add token authorization header
		header, errHeader := middleware.MakeAuthenticatedRequest(tokenGenerate)
		helpers.LoggerIfError(errHeader)
		log.Info("Add to header authorization: ", header.Body.Close())

		optionalUserLoginHistory := helpers.NewOptionalColumnParams{
			Token:   tokenGenerate,
			TimeAt:  time.Now(),
			TimeOut: time.Time{},
		}

		userLoginHistory := helpers.UserAndAccountRecordToAccountLoginHistoryRecord(
			userRecord, accountRecord, optionalUserLoginHistory)

		log.Info("User login histories record: ", userLoginHistory)

		saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
		helpers.LoggerIfError(saveLoginHistory)

		userAccounts := helpers.RecordToAccountDTO(credentialAccount)

		return userAccounts, optionalUserLoginHistory.Token, nil
	}
}
