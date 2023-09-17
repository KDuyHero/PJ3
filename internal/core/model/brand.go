package coreModel

import "mobile-ecommerce/internal/common"

type BrandModel struct {
	Name        string
	Description string
	Slug        string
}
type BrandsModel struct {
	Brands     []BrandModel
	Pagination common.Pagination
}
