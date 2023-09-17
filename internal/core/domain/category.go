package domain

import (
	"context"
	"mobile-ecommerce/internal/common"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

type GetCategoriesRepoParams struct {
	common.GetQueryFields
}

type CreateCategoryRepoParams struct {
	Name string
	Slug string
}

type CategoryUpdateInfo struct {
	Name string
}

type UpdateCategoryBySlugRepoParams struct {
	Slug    string
	NewInfo CategoryUpdateInfo
}
type CategoryRepository interface {
	GetCategories(ctx context.Context, params GetCategoriesRepoParams) ([]coreEntity.Category, *common.Pagination, error)
	CreateNewCategory(ctx context.Context, params CreateCategoryRepoParams) (string, error)
	GetCategoryByName(ctx context.Context, name string) (*coreEntity.Category, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*coreEntity.Category, error)
	DeleteCategoryBySlug(ctx context.Context, slug string) error
	UpdateCategoryBySlug(ctx context.Context, newInfo UpdateCategoryBySlugRepoParams) error
}

type GetCategoriesUcParams struct {
	UserId int64
	common.GetQueryFields
}
type AddCategoryUcParams struct {
	UserId int64
	Name   string
}
type UpdateCategoryBySlugUcParams struct {
	UserId int64
	Slug   string
	Name   string
}
type CategoryUsecase interface {
	GetCategories(ctx context.Context, params GetCategoriesUcParams) (*coreModel.CategoriesModel, error)
	AddNewCategory(ctx context.Context, info AddCategoryUcParams) error
	GetCategoryBySlug(ctx context.Context, userId int64, slug string) (*coreModel.CategoryModel, error)
	DeleteCategoryBySlug(ctx context.Context, userId int64, slug string) error
	UpdateCategoryBySlug(ctx context.Context, params UpdateCategoryBySlugUcParams) error
}
