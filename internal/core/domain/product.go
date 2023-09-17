package domain

import (
	"context"
	"mobile-ecommerce/internal/common"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

type GetProductsRepoParams struct {
	common.GetQueryFields
}
type CreateProductRepoParams struct {
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}

type UpdateProductInfo struct {
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}
type UpdateProductByIdRepoParams struct {
	ProductId int64
	NewInfo   UpdateProductInfo
}
type ProductRepository interface {
	GetProducts(ctx context.Context, params GetProductsRepoParams) (*[]coreEntity.Product, *common.Pagination, error)
	GetProductById(ctx context.Context, productId int64) (*coreEntity.Product, error)
	CreateProduct(ctx context.Context, params CreateProductRepoParams) (int64, error)
	UpdateProductById(ctx context.Context, params UpdateProductByIdRepoParams) error
	DeleteProductById(ctx context.Context, productId int64) error
}

type GetProductsUcParams struct {
	common.GetQueryFields
}

type AddProductUcParams struct {
	UserId     int64
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}

type UpdateProductParams struct {
	UserId     int64
	ProductId  int64
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}
type ProductUsecase interface {
	GetProducts(ctx context.Context, params GetProductsUcParams) (*coreModel.ProductsModel, error)
	GetProductById(ctx context.Context, productId int64) (*coreModel.ProductModel, error)
	AddProduct(ctx context.Context, params AddProductUcParams) error
	UpdateProductById(ctx context.Context, params UpdateProductParams) error
	DeleteProductById(ctx context.Context, userId int64, productId int64) error
}
