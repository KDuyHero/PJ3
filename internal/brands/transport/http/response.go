package brand_http

import coreModel "mobile-ecommerce/internal/core/model"

type getBrandsResponse struct {
	Code int
	Data coreModel.BrandsModel
}

func newGetBrandsResponse(brandsModel coreModel.BrandsModel) getBrandsResponse {
	return getBrandsResponse{
		Code: 0,
		Data: brandsModel,
	}
}

type getBrandResponse struct {
	Code int
	Data coreModel.BrandModel
}

func newGetBrandResponse(brandModel coreModel.BrandModel) getBrandResponse {
	return getBrandResponse{
		Code: 0,
		Data: brandModel,
	}
}
