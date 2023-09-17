package product_usecase

import (
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

func MapProductEntityToCoreModel(product coreEntity.Product) coreModel.ProductModel {
	return coreModel.ProductModel{
		Id:         product.Id,
		Name:       product.Name,
		BrandName:  product.BrandName,
		Thumbnail:  product.Thumbnail,
		Properties: product.Properties,
	}
}

func MapListProductsEntityToCoreModel(listProducts []coreEntity.Product) []coreModel.ProductModel {
	var products []coreModel.ProductModel
	for _, product := range listProducts {
		products = append(products, MapProductEntityToCoreModel(product))
	}
	return products
}
