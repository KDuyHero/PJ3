package brandRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapBrandGormToEntity(brand gormModel.Brand) coreEntity.Brand {
	return coreEntity.Brand{
		Name:        brand.Name,
		Description: brand.Description.String,
		Slug:        brand.Slug,
	}
}

func MapListBrandsGormToEntity(listBrands []gormModel.Brand) []coreEntity.Brand {
	var brands []coreEntity.Brand
	for _, brand := range listBrands {
		brands = append(brands, MapBrandGormToEntity(brand))
	}
	return brands
}
