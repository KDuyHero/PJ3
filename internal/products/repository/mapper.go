package productRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapProductGormModelToEntity(product gormModel.Product) coreEntity.Product {

	listPro := []struct {
		Name  string
		Value string
	}{}
	for _, pro := range product.Properties {
		listPro = append(listPro, struct {
			Name  string
			Value string
		}{
			Name:  pro.Name,
			Value: pro.Value,
		})
	}

	return coreEntity.Product{
		Id:         product.Id,
		Name:       product.Name,
		Thumbnail:  product.Thumbnail.String,
		BrandName:  product.BrandName,
		Properties: listPro,
	}
}

func MapListProductsGormToEntity(listProducts []gormModel.Product) []coreEntity.Product {
	var products []coreEntity.Product
	for _, product := range listProducts {
		products = append(products, MapProductGormModelToEntity(product))
	}

	return products
}
