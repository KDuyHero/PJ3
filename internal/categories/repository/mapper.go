package categoryRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapCategoryGormModelToEntity(category gormModel.Category) coreEntity.Category {
	return coreEntity.Category{
		Name: category.Name,
		Slug: category.Slug,
	}
}

func MapListCategoriesGormModelToEntity(listCategories []gormModel.Category) []coreEntity.Category {
	var categories []coreEntity.Category
	for _, category := range listCategories {
		categories = append(categories, MapCategoryGormModelToEntity(category))
	}
	return categories
}
