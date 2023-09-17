package user_http

import (
	"fmt"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	"mobile-ecommerce/internal/util"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(engine *gin.RouterGroup, userUsecase domain.UserUsecase, tokenMaker token.TokenMaker) {
	handler := &userHandler{
		userUsecase: userUsecase,
	}
	engine.GET("/is-admin", middleware.RequireAuthMiddleware(tokenMaker), handler.isAdmin)
	group := engine.Group("/users")
	group.GET("", middleware.RequireAuthMiddleware(tokenMaker), handler.GetUsers)
	group.GET("/:id", middleware.RequireAuthMiddleware(tokenMaker), handler.GetUserInfoById)
	group.PUT("/:id", middleware.RequireAuthMiddleware(tokenMaker), handler.UpdateUserInfo)

}

// Get list users
// @Name GetUsers
// @Description Get list of users with conditon
// @Tags Users
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param request query GetUsersRequest true "query params"
// @Success 200 {object} getUsersResponse "Success"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 401 {object} common.BaseResponse "Unathorization"
// @Failure 403 {object} common.BaseResponse "Forbidden"
// @Failure 500 {object} common.BaseResponse "Unexpected error"
// @Router /users [get]
func (handler *userHandler) GetUsers(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
	}
	request := GetUsersRequest{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
	}

	userUcParams := domain.GetUsersUsecaseParams{
		KeyWord: request.KeyWord,
		Limit:   request.Limit,
		Page:    request.Page,
		OrderBy: request.OrderBy,
		Role:    userInfo.Role,
	}

	listUsersModel, err := handler.userUsecase.GetUsers(userUcParams)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetUsersResponse(*listUsersModel))
}

// Get user info by id
// @Name Get user by id
// @Description Get user by id
// @Tags Users
// @Produce json
// @Param authorization header string true "access token"
// @Param id path int64 true "User id"
// @Success 200 {object} getUserByIdResponse "Success"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 401 {object} common.BaseResponse "Unauthorization"
// @Failure 403 {object} common.BaseResponse "Forbidden"
// @Failure 500 {object} common.BaseResponse "Unexpected error"
// @Router /users/{id} [get]
func (handler *userHandler) GetUserInfoById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	fmt.Println(idParam)
	idUser, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, "id must be int"))
		return
	}

	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	user, err := handler.userUsecase.GetUserById(domain.GetUserByIdUsecaseParams{
		Ctx:           ctx,
		IdUser:        idUser,
		IdUserRequest: userInfo.UserId,
	})

	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, NewGetUserByIdResponse(*user))
}

// Update user info
// @Name Update user info
// @Description Update user info by id
// @Tags Users
// @Produce json
// @Param authorization header string true "access token"
// @Param id path int64 true "User id"
// @Success 200 {object} common.BaseResponse "Success"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 401 {object} common.BaseResponse "Unauthorization"
// @Failure 403 {object} common.BaseResponse "Forbidden"
// @Failure 500 {object} common.BaseResponse "Unexpected error"
// @Router /users/{id} [get]
func (handler *userHandler) UpdateUserInfo(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUserUpdate, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, "id must be int"))
		return
	}
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	updateUserBody := domain.UpdateUserByIdUsecaseParams{}
	err = ctx.ShouldBindJSON(&updateUserBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}
	err = handler.userUsecase.UpdateUserById(domain.UpdateUserByIdUsecaseParams{
		Ctx:             ctx,
		IdUserUpdate:    idUserUpdate,
		IdUserRequest:   userInfo.UserId,
		RoleUserRequest: userInfo.Role,
		UserName:        updateUserBody.UserName,
		Name:            updateUserBody.Name,
		PhoneNumber:     updateUserBody.PhoneNumber,
		Avatar:          updateUserBody.Avatar,
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}

func (handler *userHandler) isAdmin(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok || !util.CheckRole(userInfo.Role, domain.RequireAdmin) {
		ctx.JSON(http.StatusForbidden, "You don't have permision")
	}

	ctx.JSON(200, gin.H{
		"ok": true,
	})

}
