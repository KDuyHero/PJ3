package brand_usecase

import (
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

func mapBrandEntityToModel(brand coreEntity.Brand) coreModel.BrandModel {
	return coreModel.BrandModel{
		Name:        brand.Name,
		Description: brand.Description,
		Slug:        brand.Slug,
	}
}

func mapListBrandsEntityToModel(listBrands []coreEntity.Brand) []coreModel.BrandModel {
	var brands []coreModel.BrandModel
	for _, brand := range listBrands {
		brands = append(brands, mapBrandEntityToModel(brand))
	}
	return brands
}
