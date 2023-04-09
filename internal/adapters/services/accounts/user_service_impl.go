package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/helpers"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
)

type UserServiceImpl struct {
	UserDetailRepository accounts.UserDetailRecordRepository
	AccountRepository    accounts.AccountRecordRepository
	Validate             *validator.Validate
}

func NewUserDetailServiceImpl(
	userDetailRepository accounts.UserDetailRecordRepository,
	accountRepository accounts.AccountRecordRepository,
	validate *validator.Validate) account.UserService {

	return &UserServiceImpl{
		UserDetailRepository: userDetailRepository,
		AccountRepository:    accountRepository,
		Validate:             validate}
}

func (u UserServiceImpl) FindUserByUserIdService(ctx context.Context, userId int64) (dto.UserDetailDTO, error) {
	log := helpers.LoggerParent()

	err := helpers.ValidateIntValue(int(userId))
	if err != nil {
		log.Warn("validate int is not valid: ", err.Error())
	}

	findUserDetail, err := u.UserDetailRepository.FindUserById(ctx, userId)
	helpers.LoggerIfError(err)

	findUserDetailResponse := helpers.UserDetailRecordToUserDetailDTO(findUserDetail)

	return findUserDetailResponse, nil
}

func (u UserServiceImpl) UpdateUserByUserIdService(ctx context.Context, userId int64) (dto.UserDetailDTO, error) {
	//TODO implement me
	panic("implement me")
}
