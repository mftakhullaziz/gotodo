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

type RegisterServiceImpl struct {
	AccountRepository accounts.AccountRecordRepository
	UserRepository    accounts.UserDetailRecordRepository
	Validate          *validator.Validate
}

func NewRegisterServiceImpl(
	accountRepository accounts.AccountRecordRepository,
	userRepository accounts.UserDetailRecordRepository,
	validate *validator.Validate) account.RegisterService {

	return &RegisterServiceImpl{
		AccountRepository: accountRepository,
		UserRepository:    userRepository,
		Validate:          validate}
}

func (r RegisterServiceImpl) CreateNewAccount(ctx context.Context, request request.RegisterRequest) (dto.AccountDTO, error) {
	log := helpers.LoggerParent()

	err := r.Validate.Struct(request)
	helpers.PanicIfError(err)

	existingUsername := r.AccountRepository.IsExistUsername(ctx, request.Username)
	existingEmail := r.AccountRepository.IsExistAccountEmail(ctx, request.Email)

	if existingEmail == true && existingUsername == true {
		log.Info("Email already registered")
		return dto.AccountDTO{}, nil
	} else {
		hashPassword := helpers.HashPasswordAndSalt([]byte(request.Password))
		userCreate := dto.UserDetailDTO{
			Username:  request.Username,
			Password:  hashPassword,
			Email:     request.Email,
			Status:    "active",
			CreatedAt: time.Now(),
		}
		userRecord := helpers.UserDTOToRecord(userCreate)
		createUser, err := r.UserRepository.SaveUser(ctx, userRecord)
		helpers.PanicIfError(err)

		accountService := dto.AccountDTO{
			UserID:    int(createUser.UserID),
			Username:  createUser.Username,
			Password:  createUser.Password,
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		accountRecord := helpers.AccountDtoToRecord(accountService)
		createAccount, err := r.AccountRepository.SaveAccount(ctx, accountRecord)
		helpers.PanicIfError(err)

		accountDTO := helpers.RecordToAccountDTO(createAccount)
		log.Info("Account created successfully")
		return accountDTO, nil
	}

}
