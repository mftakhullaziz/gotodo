package accounts

import (
	"gotodo/internal/domain/models/request"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
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
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		responses := utils.BuildEmptyResponse(messageUserNotAuthorized)
		// Do build write response to response body
		utils.WriteToResponseBody(writer, &responses)
		return
	}

	authorizedUserId, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	findUserDetailHandler, err := u.UserUsecase.FindUserByUserIdUsecase(requests.Context(), int64(authorizedUserId))

	responses := utils.CreateResponses(findUserDetailHandler, http.StatusOK,
		"find user detail successfully!",
		"find user detail not success please check your username or email or password again!",
	)

	utils.WriteToResponseBody(writer, &responses)
}

func (u UserDetailHandlerAPI) UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request) {
	log := utils.LoggerParent()

	token := requests.Header.Get(authHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	if authorized == "" {
		responses := utils.BuildEmptyResponse(messageUserNotAuthorized)
		utils.WriteToResponseBody(writer, &responses)
		return
	}

	userIsAuthorize, err := strconv.Atoi(authorized)
	utils.LoggerIfError(err)

	userUpdateRequest := request.UserRequest{}
	utils.ReadFromRequestBody(requests, &userUpdateRequest)
	log.Infoln("update user request: ", userUpdateRequest)

	updateUserDetailHandler, err := u.UserUsecase.UpdateUserByUserIdUsecase(requests.Context(), int64(userIsAuthorize), userUpdateRequest)

	updateUserDetailHandlerRes := utils.BuildResponseWithAuthorization(
		updateUserDetailHandler,
		http.StatusCreated,
		int(updateUserDetailHandler.UserID),
		authorized,
		"",
	)

	utils.WriteToResponseBody(writer, &updateUserDetailHandlerRes)
}

func (u UserDetailHandlerAPI) DeleteUserHandler(writer http.ResponseWriter, requests *http.Request) {
	//TODO implement me
	panic("implement me")
}
