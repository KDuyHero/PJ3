package cart_roduct_handler

import (
	"mobile-ecommerce/internal/core/domain"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type cartProductHandler struct {
	cartProductUsecase domain.CartProductUsecase
}

func NewCartProductHandler(engine *gin.RouterGroup, cartProductUsecase domain.CartProductUsecase, tokenMaker token.TokenMaker) {
	handler := &cartProductHandler{
		cartProductUsecase: cartProductUsecase,
	}

	group := engine.Group("/carts-products")
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddProductToCart)
	group.PUT("/", middleware.RequireAuthMiddleware(tokenMaker), handler.UpdateQuantity)
	group.DELETE("/remove", middleware.RequireAuthMiddleware(tokenMaker), handler.RemoveProduct)
	group.DELETE("/clear", middleware.RequireAuthMiddleware(tokenMaker), handler.ClearCart)

}

// Get Cart for user
// @Name AddProductToCart
// @Description add product to cart
// @Tags Carts-Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /carts-products [post]
func (handler *cartProductHandler) AddProductToCart(ctx *gin.Context) {}

// Change quantity
// @Name ChangeQuantity
// @Description change quantity of product in cart
// @Tags Carts-Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /carts-products [put]
func (handler *cartProductHandler) UpdateQuantity(ctx *gin.Context) {}

// Get Cart for user
// @Name RemoveProduct
// @Description remove specific product in cart
// @Tags Carts-Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /carts-products/remove [delete]
func (handler *cartProductHandler) RemoveProduct(ctx *gin.Context) {}

// Get Cart for user
// @Name ClearCart
// @Description remove all product in cart
// @Tags Carts-Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /carts-products/clear [delete]
func (handler *cartProductHandler) ClearCart(ctx *gin.Context) {}
