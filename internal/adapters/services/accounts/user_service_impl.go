package accounts

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gotodo/internal/domain/dto"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/ports/repositories/accounts"
	account "gotodo/internal/ports/services/accounts"
	"gotodo/internal/utils/converter"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"gotodo/internal/utils/password"
	validate "gotodo/internal/utils/validator"
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
	log := logger.LoggerParent()

	err := validate.ValidateIntValue(int(userId))
	if err != nil {
		log.Warn("validate int is not valid: ", err.Error())
	}

	findUserDetail, err := u.UserDetailRepository.FindUserById(ctx, userId)
	errs.LoggerIfError(err)

	findUserDetailResponse := converter.UserDetailRecordToUserDetailDTO(findUserDetail)

	return findUserDetailResponse, nil
}

func (u UserServiceImpl) UpdateUserByUserIdService(ctx context.Context, userId int64, request request.UserRequest) (dto.UserDetailDTO, error) {
	log := logger.LoggerParent()

	err := u.Validate.Struct(request)
	if err != nil {
		log.Warn("validate update request is errors: ", err.Error())
	}
	hashPassword := password.HashPasswordAndSalt([]byte(request.Password))

	userUpdate := dto.UserDetailDTO{
		Email:       request.Email,
		Password:    hashPassword,
		Name:        request.Name,
		MobilePhone: request.MobilePhone,
		Address:     request.Address,
		Status:      request.Status,
	}

	userRecord := converter.UserDTOToRecord(userUpdate)
	updateUser, err := u.UserDetailRepository.UpdateUser(ctx, userId, userRecord)
	errs.LoggerIfError(err)

	accountService := dto.AccountDTO{
		Password: updateUser.Password,
		Status:   updateUser.Status,
	}

	accountRecord := converter.AccountDtoToRecord(accountService)
	accountUser, err := u.AccountRepository.UpdateAccount(ctx, userId, accountRecord)
	log.Infoln("account is updated: ", accountUser)

	userUpdateResponse := converter.RecordToUserDTO(updateUser)
	return userUpdateResponse, nil
}
