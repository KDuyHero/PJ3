package category_http

import coreModel "mobile-ecommerce/internal/core/model"

type getCategoriesResponse struct {
	Code int
	Data coreModel.CategoriesModel
}

func newGetCategoriesResponse(model coreModel.CategoriesModel) getCategoriesResponse {
	return getCategoriesResponse{
		Code: 0,
		Data: model,
	}
}

type getCategoryResponse struct {
	Code     int
	Category coreModel.CategoryModel
}

func newGetCategoryResponse(category coreModel.CategoryModel) getCategoryResponse {
	return getCategoryResponse{
		Code:     0,
		Category: category,
	}
}
