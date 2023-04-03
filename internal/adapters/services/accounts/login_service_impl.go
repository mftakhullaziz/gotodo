package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
	"time"
)

type LoginServiceImpl struct {
	AccountRepository accounts.AccountRecordRepository
	Validate          *validator.Validate
}

func NewLoginServiceImpl(accountRepository accounts.AccountRecordRepository, validate *validator.Validate) account.LoginService {
	return &LoginServiceImpl{AccountRepository: accountRepository, Validate: validate}
}

func (l LoginServiceImpl) VerifyCredentialAccount(ctx context.Context, request request.LoginRequest) (dto.AccountDTO, error) {
	log := helpers.LoggerParent()

	validate := l.Validate.Struct(request)
	helpers.PanicIfError(validate)

	credentialAccount, err := l.AccountRepository.VerifyCredential(ctx, request.Username)
	helpers.PanicIfErrorWithCustomMessage(err, "Username not found")

	comparedPassword, errCompare := helpers.ComparedPassword(credentialAccount.Password, []byte(request.Password))
	if comparedPassword != true && errCompare != nil {
		log.Info("Password not matched")
		return dto.AccountDTO{}, errCompare
	} else {
		log.Info("Password is matched")

		findUserAccount, errUserAccount := l.AccountRepository.FindAccountUser(ctx, credentialAccount.Username)
		if errUserAccount != nil {
			log.Error("Error find user account not found", errUserAccount.Error())
		}
		log.Info("User Account Find: ", findUserAccount)

		userRecord := findUserAccount.Users
		accountRecord := findUserAccount.Accounts

		userLoginHistory := helpers.UserAndAccountRecordToAccountLoginHistoryRecord(
			userRecord, accountRecord, "active", time.Now(), time.Time{})

		log.Info("User login histories record: ", userLoginHistory)

		saveLoginHistory := l.AccountRepository.SaveLoginHistories(ctx, userLoginHistory)
		helpers.LoggerIfError(saveLoginHistory)

		userAccounts := helpers.RecordToAccountDTO(credentialAccount)

		return userAccounts, nil
	}
}
