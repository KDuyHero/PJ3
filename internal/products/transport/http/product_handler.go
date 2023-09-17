package product_http

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

type productHandler struct {
	productUc domain.ProductUsecase
}

func NewProductHandler(engine *gin.RouterGroup, productUc domain.ProductUsecase, tokenMaker token.TokenMaker) {
	handler := &productHandler{
		productUc: productUc,
	}

	group := engine.Group("products")
	group.GET("", handler.GetProducts)
	group.GET("/:id", handler.GetProduct)
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddProduct)
	group.PUT("/:id", middleware.RequireAuthMiddleware(tokenMaker), handler.UpdateProduct)
	group.DELETE("/:id", middleware.RequireAuthMiddleware(tokenMaker), handler.DeleteProduct)

}

// Get List Products
// @Name GetProducts
// @Description Get list products with condition
// @Tags Products
// @Produce json
// @Param request query getProductsRequest false "query params"
// @Success 200 {object} getProductsResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /products [get]
func (handler *productHandler) GetProducts(ctx *gin.Context) {
	request := getProductsRequest{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	productsModel, err := handler.productUc.GetProducts(ctx, domain.GetProductsUcParams{
		GetQueryFields: common.GetQueryFields{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Page:    request.Page,
		},
	})

	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetProductsResponse(*productsModel))
}

// Get A Product
// @Name GetProduct
// @Description Get a product with Its id
// @Tags Products
// @Produce json
// @Param id path int64 true "product's id"
// @Success 200 {object} getProductResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 404 {object} common.BaseResponse "product not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /products/{id} [get]
func (handler *productHandler) GetProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	productId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusAccepted, common.NewErrResponse(http.StatusBadRequest, "product's id must be int"))
		return
	}

	productModel, err := handler.productUc.GetProductById(ctx, productId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetProductResponse(*productModel))

}

// Add new product
// @Name AddProduct
// @Description Add a product
// @Tags Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Param request body addProductRequest true "new product info"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /products [post]
func (handler *productHandler) AddProduct(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	body := addProductRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.productUc.AddProduct(ctx, domain.AddProductUcParams{
		UserId:    userInfo.UserId,
		Name:      body.Name,
		Thumbnail: body.Thumbnail,
		BrandName: body.BrandName,
		Properties: []struct {
			Name  string
			Value string
		}(body.Properties),
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}

// Update product
// @Name UpdateProduct
// @Description Update a product
// @Tags Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Param id path int64 true "product's id"
// @Param request body updateProductRequest true "new info"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 404 {object} common.BaseResponse "product not found"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /products/{id} [put]
func (handler *productHandler) UpdateProduct(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	idParam := ctx.Param("id")
	productId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusAccepted, common.NewErrResponse(http.StatusBadRequest, "product's id must be int"))
		return
	}

	body := updateProductRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	fmt.Println(body)
	err = handler.productUc.UpdateProductById(ctx, domain.UpdateProductParams{
		UserId:    userInfo.UserId,
		ProductId: productId,
		Name:      body.Name,
		Thumbnail: body.Thumbnail,
		BrandName: body.BrandName,
		Properties: []struct {
			Name  string
			Value string
		}(body.Properties),
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}

// Delete product
// @Name DeleteProduct
// @Description Delete a product
// @Tags Products
// @Produce json
// @Param authorization header string true "bearer token"
// @Param id path int64 true "product's id"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /products/{id} [delete]
func (handler *productHandler) DeleteProduct(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	idParam := ctx.Param("id")
	productId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusAccepted, common.NewErrResponse(http.StatusBadRequest, "product's id must be int"))
		return
	}

	err = handler.productUc.DeleteProductById(ctx, userInfo.UserId, productId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}
