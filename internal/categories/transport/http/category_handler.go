package category_http

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

type categoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(engine *gin.RouterGroup, categoryUsecase domain.CategoryUsecase, tokenMaker token.TokenMaker) {
	handler := &categoryHandler{
		categoryUsecase: categoryUsecase,
	}

	group := engine.Group("categories")
	group.GET("", middleware.RequireAuthMiddleware(tokenMaker), handler.GetCategories)
	group.GET("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.GetCategoryBySlug)
	group.POST("", middleware.RequireAuthMiddleware(tokenMaker), handler.AddNewCategory)
	group.PUT("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.UpdateCategoryBySlug)
	group.DELETE("/:slug", middleware.RequireAuthMiddleware(tokenMaker), handler.DeleteCategoryBySlug)

}

// get categories
// @Name GetCategories
// @Description Get list categories
// @Tags Categories
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param request query getCategoriesRequest fasle "query params"
// @Success 200 {object} getCategoriesResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /categories [get]
func (handler *categoryHandler) GetCategories(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	query := getCategoriesRequest{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	categoriesModel, err := handler.categoryUsecase.GetCategories(ctx, domain.GetCategoriesUcParams{
		UserId: userInfo.UserId,
		GetQueryFields: common.GetQueryFields{
			Keyword: query.Keyword,
			Page:    query.Page,
			Limit:   query.Limit,
		},
	})

	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetCategoriesResponse(*categoriesModel))
}

// get a category
// @Name GetCategory
// @Description Get a category by slug
// @Tags Categories
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "category slug"
// @Success 200 {object} getCategoryResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /categories/{slug} [get]
func (handler *categoryHandler) GetCategoryBySlug(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	nameParam := ctx.Param("slug")

	category, err := handler.categoryUsecase.GetCategoryBySlug(ctx, userInfo.UserId, nameParam)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, newGetCategoryResponse(*category))
}

// Delete a category
// @Name DeleteCategory
// @Description Delete a category by slug
// @Tags Categories
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "category slug"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /categories/{slug} [delete]
func (handler *categoryHandler) DeleteCategoryBySlug(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	nameParam := ctx.Param("slug")
	err := handler.categoryUsecase.DeleteCategoryBySlug(ctx, userInfo.UserId, nameParam)
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}

// Update a category
// @Name UpdateCategory
// @Description Update a category by slug
// @Tags Categories
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param slug path string true "category slug"
// @Param request body updateCategoryRequest true "new info"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /categories/{slug} [put]
func (handler *categoryHandler) UpdateCategoryBySlug(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	slug := ctx.Param("slug")
	body := updateCategoryRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.categoryUsecase.UpdateCategoryBySlug(ctx, domain.UpdateCategoryBySlugUcParams{
		UserId: userInfo.UserId,
		Slug:   slug,
		Name:   body.Name,
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())

}

// Create a category
// @Name CreateCategory
// @Description Create a category by slug
// @Tags Categories
// @Produce json
// @Param authorization header string true "Bearer token"
// @Param request body addCategoryRequest true "new info"
// @Success 200 {object} common.BaseResponse "success"
// @Failure 400 {object} common.BaseResponse "invalid request"
// @Failure 401 {object} common.BaseResponse "unauthorization"
// @Failure 403 {object} common.BaseResponse "forbidden"
// @Failure 500 {object} common.BaseResponse "unexpected error"
// @Router /categories [post]
func (handler *categoryHandler) AddNewCategory(ctx *gin.Context) {
	userInfo, ok := util.GetUserInfoFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, common.NewUnAuthorizeResponse())
		return
	}

	body := addCategoryRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewErrResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.categoryUsecase.AddNewCategory(ctx, domain.AddCategoryUcParams{
		UserId: userInfo.UserId,
		Name:   body.Name,
	})
	if err != nil {
		response := coreError.MapCoreErrToResponse(err)
		ctx.JSON(coreError.GetHttpStatusFromCoreErrResponse(response), response)
		return
	}

	ctx.JSON(http.StatusOK, common.NewSuccessResponse())
}
