package coreModel

import "mobile-ecommerce/internal/common"

type CategoryModel struct {
	Name string
	Slug string
}

type CategoriesModel struct {
	Categories []CategoryModel
	Pagination common.Pagination
}
