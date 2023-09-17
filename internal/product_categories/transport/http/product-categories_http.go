package product_categories_http

import (
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type productCategoriesHandler struct {
}

func NewProductCategoriesHandler(engine *gin.RouterGroup, tokenMaker token.TokenMaker) {
	handler := &productCategoriesHandler{}

	group := engine.Group("/product-categories")
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddCategoryForProduct)
	group.DELETE("/remove", middleware.RequireAuthMiddleware(tokenMaker), handler.RemoveCategory)
	group.DELETE("/clear", middleware.RequireAuthMiddleware(tokenMaker), handler.ClearProductCategories)
}

// Get Cart for user
// @Name AddCategoryForProduct
// @Description add product to cart
// @Tags Product-Categories
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /product-categories [post]
func (handler *productCategoriesHandler) AddCategoryForProduct(ctx *gin.Context) {}

// Get Cart for user
// @Name RemoveCategory
// @Description remove specific product in cart
// @Tags Product-Categories
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /product-categories/remove [delete]
func (handler *productCategoriesHandler) RemoveCategory(ctx *gin.Context) {}

// Get Cart for user
// @Name ClearProductCategories
// @Description remove all product in cart
// @Tags Product-Categories
// @Produce json
// @Param authorization header string true "bearer token"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 404 {object} common.BaseResponse "not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /product-categories/clear [delete]
func (handler *productCategoriesHandler) ClearProductCategories(ctx *gin.Context) {}
