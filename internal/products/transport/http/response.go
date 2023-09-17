package product_http

import coreModel "mobile-ecommerce/internal/core/model"

type getProductResponse struct {
	Code    int
	Product coreModel.ProductModel
}

func newGetProductResponse(productModel coreModel.ProductModel) getProductResponse {
	return getProductResponse{
		Code:    0,
		Product: productModel,
	}
}

type getProductsResponse struct {
	Code int
	Data coreModel.ProductsModel
}

func newGetProductsResponse(productsModel coreModel.ProductsModel) getProductsResponse {
	return getProductsResponse{
		Code: 0,
		Data: productsModel,
	}
}
