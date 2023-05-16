package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	errs "gotodo/internal/utils/errors"
	"gotodo/internal/utils/logger"
	"gotodo/internal/utils/payload"
	"gotodo/internal/utils/response"
	"net/http"
	"strconv"
)

const (
	// formatDatetime is the format string for datetime values.
	formatDatetime = "2006-01-02 15:04:05"
	// messageUserNotAuthorized is the errors message for unauthorized user.
	messageUserNotAuthorized = "user account not authorized, please login or sign up!"
	// authHeaderKey is the key for the Authorization header.
	authHeaderKey = "Authorization"
)

type UserDetailHandlerAPI struct {
	UserUsecase accounts.UserDetailUsecase
}

func NewUserDetailHandlerAPI(userUsecase accounts.UserDetailUsecase) api.UserHandlerAPI {
	return &UserDetailHandlerAPI{UserUsecase: userUsecase}
}

func (u UserDetailHandlerAPI) FindDataUserDetailHandler(writer http.ResponseWriter, requests *http.Request) {
	// Do get authorization token if any from user login
	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		responses := response.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		payload.WriteToResponseBody(writer, &responses)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	findUserDetailHandler, err := u.UserUsecase.FindUserByUserIdUsecase(requests.Context(), int64(authorizedUserId))

	responses := response.CreateResponses(findUserDetailHandler, http.StatusOK,
		"find user detail successfully!",
		"find user detail not success please check your username or email or password again!",
	)

	payload.WriteToResponseBody(writer, &responses)
}

func (u UserDetailHandlerAPI) UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request) {
	log := logger.LoggerParent()

	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	errs.LoggerIfError(err)

	if authorized == "" {
		responses := response.BuildEmptyResponse(messageUserNotAuthorized)
		payload.WriteToResponseBody(writer, &responses)
		return
	}

	userIsAuthorize, err := strconv.Atoi(authorized)
	errs.LoggerIfError(err)

	userUpdateRequest := request.UserRequest{}
	payload.ReadFromRequestBody(requests, &userUpdateRequest)
	log.Infoln("update user request: ", userUpdateRequest)

	updateUserDetailHandler, err := u.UserUsecase.UpdateUserByUserIdUsecase(requests.Context(), int64(userIsAuthorize), userUpdateRequest)

	updateUserDetailHandlerRes := response.BuildResponseWithAuthorization(
		updateUserDetailHandler,
		http.StatusCreated,
		int(updateUserDetailHandler.UserID),
		authorized,
		"",
	)

	payload.WriteToResponseBody(writer, &updateUserDetailHandlerRes)
}

func (u UserDetailHandlerAPI) DeleteUserHandler(writer http.ResponseWriter, requests *http.Request) {
	//TODO implement me
	panic("implement me")
}
