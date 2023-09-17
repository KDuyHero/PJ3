package cart_http

import (
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	"mobile-ecommerce/internal/util"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartUsecase domain.CartUsecase
}

func NewCartHandler(engine *gin.RouterGroup, cartUsecase domain.CartUsecase, tokenMaker token.TokenMaker) {
	handler := &cartHandler{
		cartUsecase: cartUsecase,
	}

	group := engine.Group("cart")
	group.GET("", middleware.RequireAuthMiddleware(tokenMaker), handler.GetCart)
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddCart)
	group.DELETE("", middleware.RequireAuthMiddleware(tokenMaker), handler.DeleteCart)
}

// Get Cart for user
// @Name GetCart
// @Description get cart with userid
// @Tags Cart
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} getCartResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /cart [get]
func (handler *cartHandler) GetCart(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	cart, err := handler.cartUsecase.GetCart(ctx, userInfo.UserId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, NewGetCartResponse(*cart))
}

// Add Cart
// @Name AddCart
// @Description add cart with userid
// @Tags Cart
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /cart [post]
func (handler *cartHandler) AddCart(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	err := handler.cartUsecase.CreateCart(ctx, domain.AddCartUcParams{
		UserId: userInfo.UserId,
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())

}

// Delete Cart
// @Name DeleteCart
// @Description delete cart with userid
// @Tags Cart
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /cart [delete]
func (handler *cartHandler) DeleteCart(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	err := handler.cartUsecase.DeleteCart(ctx, userInfo.UserId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())

}
