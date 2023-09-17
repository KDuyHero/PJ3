package brand_usecase

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

type brandUsecase struct {
	brandRepo domain.BrandRepository
	userRepo  domain.UserRepository
}

func NewBrandUsecase(brandRepo domain.BrandRepository, userRepo domain.UserRepository) domain.BrandUsecase {
	return &brandUsecase{
		brandRepo: brandRepo,
		userRepo:  userRepo,
	}
}

func (brandUC *brandUsecase) AddNewBrand(ctx context.Context, params domain.AddBrandUCParams) error {
	// only admin can add brand
	user, err := brandUC.userRepo.GetUserById(ctx, params.UserId)
	if err != nil {
		return coreError.ErrUserNotFound
	}
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check brand existed
	brand, _ := brandUC.brandRepo.GetBrandByName(ctx, params.Name)
	if brand != nil {
		return coreError.ErrBrandExisted
	}

	// gen slug from name (slug is unique)
	brandSlug := slug.Make(params.Name)
	brand, _ = brandUC.brandRepo.GetBrandBySlug(ctx, brandSlug)
	if brand != nil {
		x := 1
		for {
			brand, _ = brandUC.brandRepo.GetBrandBySlug(ctx, brandSlug+"-"+strconv.Itoa(x))
			if brand == nil {
				brandSlug += "-" + strconv.Itoa(x)
				break
			}
			x += 1
		}
	}

	// create brand
	_, err = brandUC.brandRepo.CreateBrand(
		ctx,
		domain.CreateBrandRepoParams{
			Name:        params.Name,
			Description: params.Description,
			Slug:        brandSlug,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (brandUC *brandUsecase) GetBrands(ctx context.Context, params domain.GetBrandsUcParams) (*coreModel.BrandsModel, error) {
	// validate pagination info
	common.CheckPagination(&params.Limit, &params.Page)
	// get brands
	brands, pagination, err := brandUC.brandRepo.GetBrands(ctx, domain.GetBrandsRepoParams{
		GetQueryFields: common.GetQueryFields{
			Keyword: params.Keyword,
			Limit:   params.Limit,
			Page:    params.Page,
		},
	})
	if err != nil {
		return nil, err
	}

	return &coreModel.BrandsModel{
		Brands:     mapListBrandsEntityToModel(brands),
		Pagination: *pagination,
	}, nil

}

func (brandUC *brandUsecase) DeleteBrandBySlug(ctx context.Context, brandSlug string, userId int64) error {
	// just admin can delete
	user, err := brandUC.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return coreError.ErrUserNotFound
	}
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check brand existed
	brand, _ := brandUC.brandRepo.GetBrandBySlug(ctx, brandSlug)
	if brand == nil {
		return coreError.ErrBrandNotFound
	}

	// if brand existed => delete
	err = brandUC.brandRepo.DeleteBrandBySlug(ctx, brandSlug)
	if err != nil {
		return err
	}
	return nil
}

func (brandUC *brandUsecase) UpdateBrandBySlug(ctx context.Context, params domain.UpdateBrandBySlugUcParams) error {
	// just admin can update
	user, err := brandUC.userRepo.GetUserById(ctx, params.UserId)
	if err != nil {
		return coreError.ErrUserNotFound
	}
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check brand existed
	brand, _ := brandUC.brandRepo.GetBrandBySlug(ctx, params.Slug)
	if brand == nil {
		return coreError.ErrBrandNotFound
	}

	// check newInfo (name)
	brand, _ = brandUC.brandRepo.GetBrandByName(ctx, params.Name)
	if brand != nil {
		return coreError.ErrBrandNameExisted
	}

	// is brand existed
	err = brandUC.brandRepo.UpdateBrandBySlug(ctx, domain.UpdateBrandBySlugRepoParams{
		Slug: params.Slug,
		NewInfo: domain.UpdateBrandInfo{
			Name:        params.Name,
			Description: params.Description,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (brandUC *brandUsecase) GetBrandBySlug(ctx context.Context, brandSlug string, userId int64) (*coreModel.BrandModel, error) {
	// just admin can get specific brand info
	user, err := brandUC.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, coreError.ErrUserNotFound
	}
	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return nil, coreError.ErrNotPermission
	}

	brand, err := brandUC.brandRepo.GetBrandBySlug(ctx, brandSlug)
	if err != nil {
		return nil, err
	}
	brandModel := mapBrandEntityToModel(*brand)
	return &brandModel, nil
}
