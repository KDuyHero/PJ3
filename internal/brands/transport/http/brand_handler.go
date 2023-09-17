package brand_http

import (
	"fmt"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	"mobile-ecommerce/internal/util"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type brandHander struct {
	brandUsecase domain.BrandUsecase
}

func NewBrandHandler(engine *gin.RouterGroup, brandUc domain.BrandUsecase, tokenMaker token.TokenMaker) {
	handler := &brandHander{
		brandUsecase: brandUc,
	}
	group := engine.Group("/brands")
	group.GET("", handler.GetBrands)
	group.GET("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.GetBrandBySlug)
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddNewBrand)
	group.PUT("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.UpdateBrandBySlug)
	group.DELETE("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.DeleteBrandBySlug)

}

// get brands
// @Name GetBrands
// @Description Get list brand with condition
// @Tags Brands
// @Produce json
// @Param request query getBrandsRequest true "query params"
// @Success 200 {object} getBrandsResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /brands [get]
func (handler *brandHander) GetBrands(ctx *gin.Context) {
	fmt.Println("get brands")
	requestBody := getBrandsRequest{}
	if err := ctx.ShouldBindQuery(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	brandsModel, err := handler.brandUsecase.GetBrands(ctx, domain.GetBrandsUcParams{
		GetQueryFields: common.GetQueryFields{
			Keyword: requestBody.KeyWord,
			Limit:   requestBody.Limit,
			Page:    requestBody.Page,
		},
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetBrandsResponse(*brandsModel))
}

// get brand by slug
// @Name GetBrandBySlug
// @Description Get a brand with slug
// @Tags Brands
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "brand's slug"
// @Success 200 {object} getBrandResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /brands/{slug} [get]
func (handler *brandHander) GetBrandBySlug(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	slug := ctx.Param("slug")
	brand, err := handler.brandUsecase.GetBrandBySlug(ctx, slug, userInfo.UserId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetBrandResponse(*brand))
}

// add new brand
// @Name AddNewBrand
// @Description Add new brand
// @Tags Brands
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param request body addNewBrandRequest true "new brand info"
// @Success 200 {object} getBrandResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /brands [post]
func (handler *brandHander) AddNewBrand(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	requestBody := addNewBrandRequest{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.brandUsecase.AddNewBrand(ctx, domain.AddBrandUCParams{
		UserId:      userInfo.UserId,
		Name:        requestBody.Name,
		Description: requestBody.Description,
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())

}

// update brand's info
// @Name UpdateBrandInfoBySlug
// @Description Update Brand info with slug param
// @Tags Brands
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "brand's slug"
// @Param newInfo body updateBrandRequest true "new info"
// @Success 200 {object} getBrandResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /brands/{slug} [put]
func (handler *brandHander) UpdateBrandBySlug(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}
	// get brand's name need to update
	slug := ctx.Param("slug")

	// get info update from body
	newInfo := updateBrandRequest{}
	if err := ctx.ShouldBindJSON(&newInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// update
	err := handler.brandUsecase.UpdateBrandBySlug(
		ctx,
		domain.UpdateBrandBySlugUcParams{
			UserId:      userInfo.UserId,
			Slug:        slug,
			Name:        newInfo.Name,
			Description: newInfo.Description,
		},
	)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}

// delete brand
// @Name DeleteBrandBySlug
// @Description Delete brand with slug
// @Tags Brands
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "brand's slug"
// @Success 200 {object} getBrandResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /brands/{slug} [delete]
func (handler *brandHander) DeleteBrandBySlug(ctx *gin.Context) {
	// get user info do delete action
	useInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	// get brand's name is deleted
	slug := ctx.Param("slug")
	// delete brand
	err := handler.brandUsecase.DeleteBrandBySlug(ctx, slug, useInfo.UserId)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}
