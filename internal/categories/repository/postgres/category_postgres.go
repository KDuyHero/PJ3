package category_postgres

import (
	"context"
	gormModel "mobile-ecommerce/db/models"
	categoryRepo "mobile-ecommerce/internal/categories/repository"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"

	"gorm.io/gorm"
)

type categoryRepository struct {
	gormDB *gorm.DB
}

func NewCategoryRepo(gormDB *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{
		gormDB: gormDB,
	}
}

func (repo *categoryRepository) GetCategories(ctx context.Context, params domain.GetCategoriesRepoParams) ([]coreEntity.Category, *common.Pagination, error) {
	var categories []gormModel.Category
	query := repo.gormDB.Model(&gormModel.Category{})
	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		query.Where("name ILIKE ?", keyword)
		query.Where("slug ILIKE ?", keyword)
	}

	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	offset := (params.Page - 1) * common.LimitPerPageDefault
	if err := query.Offset(int(offset)).Limit(int(params.Limit)).Find(&categories).Error; err != nil {
		return nil, nil, err
	}
	pagination := common.Pagination{
		Limit:     params.Limit,
		Page:      params.Page,
		TotalRows: totalRows,
	}
	return categoryRepo.MapListCategoriesGormModelToEntity(categories), &pagination, nil

}

func (repo *categoryRepository) CreateNewCategory(ctx context.Context, params domain.CreateCategoryRepoParams) (string, error) {
	category := gormModel.Category{
		Name: params.Name,
		Slug: params.Slug,
	}
	if err := repo.gormDB.Create(&category).Error; err != nil {
		return "", err
	}

	return category.Name, nil
}

func (repo *categoryRepository) GetCategoryByName(ctx context.Context, name string) (*coreEntity.Category, error) {
	category := gormModel.Category{}

	if err := repo.gormDB.Where("name = ?", name).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrCategoryNotFound
		}
		return nil, err
	}

	entity := categoryRepo.MapCategoryGormModelToEntity(category)
	return &entity, nil
}

func (repo *categoryRepository) GetCategoryBySlug(ctx context.Context, slug string) (*coreEntity.Category, error) {
	category := gormModel.Category{}

	if err := repo.gormDB.Where("slug = ?", slug).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrCategoryNotFound
		}
		return nil, err
	}

	entity := categoryRepo.MapCategoryGormModelToEntity(category)
	return &entity, nil
}

func (repo *categoryRepository) DeleteCategoryBySlug(ctx context.Context, slug string) error {
	if err := repo.gormDB.Table(gormModel.Category{}.TableName()).Where("slug = ?", slug).Delete(&gormModel.Category{}).Error; err != nil {
		return err
	}

	return nil
}
func (repo *categoryRepository) UpdateCategoryBySlug(ctx context.Context, params domain.UpdateCategoryBySlugRepoParams) error {
	category := gormModel.Category{
		Name: params.NewInfo.Name,
	}

	if err := repo.gormDB.Model(&gormModel.Category{}).Where("slug = ?", params.Slug).Updates(&category).Error; err != nil {
		return err
	}

	return nil
}
