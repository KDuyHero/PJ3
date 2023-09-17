package category_usecase

import (
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

func mapCategoryEntityToCoreModel(entity coreEntity.Category) coreModel.CategoryModel {
	return coreModel.CategoryModel{
		Name: entity.Name,
		Slug: entity.Slug,
	}
}

func mapListCategoriesEntityToCoreModel(listEntity []coreEntity.Category) []coreModel.CategoryModel {
	var categories []coreModel.CategoryModel
	for _, category := range listEntity {
		categories = append(categories, mapCategoryEntityToCoreModel(category))
	}
	return categories
}
