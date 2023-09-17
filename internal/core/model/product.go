package coreModel

import "mobile-ecommerce/internal/common"

type ProductModel struct {
	Id         int64
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}

type ProductsModel struct {
	Products   []ProductModel
	Pagination common.Pagination
}
