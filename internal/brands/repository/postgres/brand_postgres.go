package brand_postgres

import (
	"context"
	"database/sql"
	"fmt"
	gormModel "mobile-ecommerce/db/models"
	brandRepo "mobile-ecommerce/internal/brands/repository"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type brandRepository struct {
	gormDB *gorm.DB
}

func NewBrandRepository(gormDB *gorm.DB) domain.BrandRepository {
	return &brandRepository{
		gormDB: gormDB,
	}
}

func (repo *brandRepository) CreateBrand(ctx context.Context, params domain.CreateBrandRepoParams) (string, error) {
	brand := gormModel.Brand{
		Name: params.Name,
		Slug: params.Slug,
		Description: sql.NullString{
			String: params.Description,
			Valid:  len(params.Description) > 0,
		},
	}

	if err := repo.gormDB.Create(&brand).Error; err != nil {
		return "", err
	}
	return brand.Name, nil
}

func (repo *brandRepository) GetBrands(ctx context.Context, params domain.GetBrandsRepoParams) ([]coreEntity.Brand, *common.Pagination, error) {
	var brands []gormModel.Brand
	// set query condition
	query := repo.gormDB.Model(&gormModel.Brand{})
	if params.Keyword != "" {
		key := "%" + params.Keyword + "%"
		query.Where(fmt.Sprintf("name ILIKE '%s' or description ILIKE '%s'", key, key))
	}
	// get total rows
	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}
	offset := (params.Page - 1) * common.LimitPerPageDefault
	if err := query.Offset(int(offset)).Limit(int(params.Limit)).Find(&brands).Error; err != nil {
		return nil, nil, err
	}

	pagination := common.Pagination{
		Limit:     params.Limit,
		Page:      params.Page,
		TotalRows: totalRows,
	}

	return brandRepo.MapListBrandsGormToEntity(brands), &pagination, nil
}

func (repo *brandRepository) DeleteBrandBySlug(ctx context.Context, brandSlug string) error {
	if err := repo.gormDB.Where("slug = ?", brandSlug).Delete(&gormModel.Brand{}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *brandRepository) GetBrandByName(ctx context.Context, name string) (*coreEntity.Brand, error) {
	brand := gormModel.Brand{}
	err := repo.gormDB.Where("name = ?", name).First(&brand).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrBrandNotFound
		}
		return nil, err
	}
	brandEntity := brandRepo.MapBrandGormToEntity(brand)
	return &brandEntity, nil
}

func (repo *brandRepository) GetBrandBySlug(ctx context.Context, brandSlug string) (*coreEntity.Brand, error) {
	brand := gormModel.Brand{}
	err := repo.gormDB.Where("slug = ?", brandSlug).First(&brand).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrBrandNotFound
		}
		return nil, err
	}
	brandEntity := brandRepo.MapBrandGormToEntity(brand)
	return &brandEntity, nil
}

func (repo *brandRepository) UpdateBrandBySlug(ctx context.Context, params domain.UpdateBrandBySlugRepoParams) error {
	brand := gormModel.Brand{
		Name: params.NewInfo.Name,
		Description: sql.NullString{
			String: params.NewInfo.Description,
			Valid:  len(params.NewInfo.Description) > 0,
		},
	}

	err := repo.gormDB.Model(gormModel.Brand{}).Where("slug = ?", params.Slug).Updates(&brand).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
