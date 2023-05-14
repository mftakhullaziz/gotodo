package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
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

func (u UserServiceImpl) UpdateUserByUserIdService(ctx context.Context, userId int64, request request.UserRequest) (dto.UserDetailDTO, error) {
	log := helpers.LoggerParent()

	err := u.Validate.Struct(request)
	if err != nil {
		log.Warn("validate update request is error: ", err.Error())
	}
	hashPassword := helpers.HashPasswordAndSalt([]byte(request.Password))

	userUpdate := dto.UserDetailDTO{
		Email:       request.Email,
		Password:    hashPassword,
		Name:        request.Name,
		MobilePhone: request.MobilePhone,
		Address:     request.Address,
		Status:      request.Status,
	}

	userRecord := helpers.UserDTOToRecord(userUpdate)
	updateUser, err := u.UserDetailRepository.UpdateUser(ctx, userId, userRecord)
	helpers.LoggerIfError(err)

	accountService := dto.AccountDTO{
		Password: updateUser.Password,
		Status:   updateUser.Status,
	}

	accountRecord := helpers.AccountDtoToRecord(accountService)
	accountUser, err := u.AccountRepository.UpdateAccount(ctx, userId, accountRecord)
	log.Infoln("account is updated: ", accountUser)

	userUpdateResponse := helpers.RecordToUserDTO(updateUser)
	return userUpdateResponse, nil
}
