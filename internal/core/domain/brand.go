package domain

import (
	"context"
	"mobile-ecommerce/internal/common"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

type GetBrandsRepoParams struct {
	common.GetQueryFields
}
type CreateBrandRepoParams struct {
	Name        string
	Description string
	Slug        string
}

type UpdateBrandInfo struct {
	Name        string
	Description string
}

// slug is contant, just update name and description
type UpdateBrandBySlugRepoParams struct {
	Slug    string
	NewInfo UpdateBrandInfo
}

type BrandRepository interface {
	GetBrands(ctx context.Context, params GetBrandsRepoParams) ([]coreEntity.Brand, *common.Pagination, error)
	DeleteBrandBySlug(ctx context.Context, slug string) error
	CreateBrand(ctx context.Context, params CreateBrandRepoParams) (string, error)
	UpdateBrandBySlug(ctx context.Context, params UpdateBrandBySlugRepoParams) error
	GetBrandBySlug(ctx context.Context, slug string) (*coreEntity.Brand, error)
	GetBrandByName(ctx context.Context, name string) (*coreEntity.Brand, error)
}

type GetBrandsUcParams struct {
	common.GetQueryFields
}
type AddBrandUCParams struct {
	UserId      int64
	Name        string
	Description string
}

type UpdateBrandBySlugUcParams struct {
	UserId      int64
	Slug        string
	Name        string
	Description string
}
type BrandUsecase interface {
	AddNewBrand(ctx context.Context, params AddBrandUCParams) error
	GetBrands(ctx context.Context, params GetBrandsUcParams) (*coreModel.BrandsModel, error)
	DeleteBrandBySlug(ctx context.Context, slug string, userId int64) error
	UpdateBrandBySlug(ctx context.Context, params UpdateBrandBySlugUcParams) error
	GetBrandBySlug(ctx context.Context, slug string, userId int64) (*coreModel.BrandModel, error)
}
