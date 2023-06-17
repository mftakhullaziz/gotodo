package accounts

import (
	"github.com/julienschmidt/httprouter"
	"gotodo/internal/adapters/handlers"
	"gotodo/internal/domain/models/request"
	"gotodo/internal/middleware"
	"gotodo/internal/ports/handlers/api"
	"gotodo/internal/ports/usecases/accounts"
	"gotodo/internal/utils"
	"net/http"
	"strconv"
)

type UserDetailHandlerAPI struct {
	UserUsecase accounts.UserDetailUsecase
}

func NewUserDetailHandlerAPI(userUsecase accounts.UserDetailUsecase) api.UserHandlerAPI {
	return &UserDetailHandlerAPI{UserUsecase: userUsecase}
}

func (u UserDetailHandlerAPI) FindDataUserDetailHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	// Do get authorization token if any from user login
	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	// Do check if user account not authorized return empty response
	if authorized == "" {
		responses := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
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

func (u UserDetailHandlerAPI) UpdateUserDetailHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	log := utils.LoggerParent().Log

	token := requests.Header.Get(handlers.AuthHeaderKey)
	authorized, err := middleware.AuthenticateUser(token)
	utils.LoggerIfError(err)

	if authorized == "" {
		responses := utils.BuildEmptyResponse(handlers.MessageUserNotAuthorized)
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

func (u UserDetailHandlerAPI) DeleteUserHandler(writer http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
