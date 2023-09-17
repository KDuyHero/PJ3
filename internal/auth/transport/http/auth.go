package auth_http

import (
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(engine *gin.RouterGroup, authUsecase domain.AuthUsecase) {
	handler := &authHandler{
		authUsecase: authUsecase,
	}
	group := engine.Group("auth")
	group.POST("/login", handler.Login)
	group.POST("/register", handler.Register)
}

// Login
// @Name Login
// @Description Login user account
// @Tags Auth
// @Produce json
// @Param request body loginRequest true "email and password account"
// @Success 200 {object} loginSuccessResponse "success"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 401 {object} common.BaseResponse "Unauthorization"
// @Failure 404 {object} common.BaseResponse "User not found"
// @Failure 500 {object} common.BaseResponse "Unexpected error"
// @Router /auth/login [post]
func (handler *authHandler) Login(ctx *gin.Context) {
	request := loginRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	token, info, err := handler.authUsecase.Login(domain.LoginParams{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		errResponse := coreError.MapCoreErrToResponse(err)
		httpCode := coreError.GetHttpStatusFromCoreErrResponse(errResponse)
		ctx.JSON(httpCode, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, newLoginSuccessResponse(token, *info))

}

// Register new user
// @Name Register
// @Description Register new user
// @Tags Auth
// @Produce json
// @Param request body registerRequest true "user info"
// @Success 200 {object} common.BaseResponse "Success"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 500 {object} common.BaseResponse "Unexpected error"
// @Router /auth/register [post]
func (handler *authHandler) Register(ctx *gin.Context) {
	request := registerRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.authUsecase.Register(domain.RegisterParams{
		Ctx:      ctx,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		errResponse := coreError.MapCoreErrToResponse(err)
		httpCode := coreError.GetHttpStatusFromCoreErrResponse(errResponse)
		ctx.JSON(httpCode, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())

}
