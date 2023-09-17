package category_usecase

import (
	"context"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	coreModel "mobile-ecommerce/internal/core/model"
	"mobile-ecommerce/internal/util"
	"strconv"

	"github.com/gosimple/slug"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
	userRepo     domain.UserRepository
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, userRepo domain.UserRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (Uc *categoryUsecase) GetCategories(ctx context.Context, params domain.GetCategoriesUcParams) (*coreModel.CategoriesModel, error) {
	// check user existed
	user, err := Uc.userRepo.GetUserById(ctx, params.UserId)
	if err != nil {
		return nil, coreError.ErrUserNotFound
	}

	// check role
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return nil, coreError.ErrNotPermission
	}

	// validate pagination params
	common.CheckPagination(&params.Limit, &params.Page)

	// get categories
	categories, pagination, err := Uc.categoryRepo.GetCategories(ctx, domain.GetCategoriesRepoParams{
		GetQueryFields: common.GetQueryFields{
			Keyword: params.Keyword,
			Page:    params.Page,
			Limit:   params.Limit,
		},
	})

	if err != nil {
		return nil, err
	}

	return &coreModel.CategoriesModel{
		Categories: mapListCategoriesEntityToCoreModel(categories),
		Pagination: *pagination,
	}, nil
}

func (Uc *categoryUsecase) AddNewCategory(ctx context.Context, params domain.AddCategoryUcParams) error {
	// check user existed
	user, err := Uc.userRepo.GetUserById(ctx, params.UserId)
	if err != nil {
		return coreError.ErrUserNotFound
	}

	// check role
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check category existed
	if category, _ := Uc.categoryRepo.GetCategoryByName(ctx, params.Name); category != nil {
		return coreError.ErrCategoryExisted
	}

	// gen slug
	categorySlug := slug.Make(params.Name)
	category, _ := Uc.categoryRepo.GetCategoryBySlug(ctx, categorySlug)
	if category != nil {
		x := 1
		for {
			category, _ = Uc.categoryRepo.GetCategoryBySlug(ctx, categorySlug+"-"+strconv.Itoa(x))
			if category == nil {
				categorySlug += "-" + strconv.Itoa(x)
				break
			}
			x += 1
		}
	}

	_, err = Uc.categoryRepo.CreateNewCategory(ctx, domain.CreateCategoryRepoParams{
		Name: params.Name,
		Slug: categorySlug,
	})

	if err != nil {
		return err
	}

	return nil
}

func (Uc *categoryUsecase) GetCategoryBySlug(ctx context.Context, userId int64, slug string) (*coreModel.CategoryModel, error) {
	// check user existed
	user, err := Uc.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, coreError.ErrUserNotFound
	}

	// check role
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return nil, coreError.ErrNotPermission
	}

	category, err := Uc.categoryRepo.GetCategoryBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	categoryModel := mapCategoryEntityToCoreModel(*category)
	return &categoryModel, nil
}

func (Uc *categoryUsecase) DeleteCategoryBySlug(ctx context.Context, userId int64, slug string) error {
	// check user existed
	user, err := Uc.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return coreError.ErrUserNotFound
	}

	// check role
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	return Uc.categoryRepo.DeleteCategoryBySlug(ctx, slug)

}
func (Uc *categoryUsecase) UpdateCategoryBySlug(ctx context.Context, params domain.UpdateCategoryBySlugUcParams) error {
	// check user existed
	user, err := Uc.userRepo.GetUserById(ctx, params.UserId)
	if err != nil {
		return coreError.ErrUserNotFound
	}

	// check role
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check category existed
	_, err = Uc.categoryRepo.GetCategoryBySlug(ctx, params.Slug)
	if err != nil {
		return err
	}

	// check new info
	category, _ := Uc.categoryRepo.GetCategoryByName(ctx, params.Name)
	if category != nil {
		return coreError.ErrCategoryNameExisted
	}

	return Uc.categoryRepo.UpdateCategoryBySlug(ctx, domain.UpdateCategoryBySlugRepoParams{
		Slug: params.Slug,
		NewInfo: domain.CategoryUpdateInfo{
			Name: params.Name,
		},
	})
}
