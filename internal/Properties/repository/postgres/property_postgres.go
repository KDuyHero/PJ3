package property_postgres

import (
	"context"
	gormModel "mobile-ecommerce/db/models"
	propertyRepo "mobile-ecommerce/internal/Properties/repository"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"

	"gorm.io/gorm"
)

type propertyRepository struct {
	gormDB *gorm.DB
}

func NewPropertyRepository(gormDB *gorm.DB) domain.PropertyRepository {
	return &propertyRepository{
		gormDB: gormDB,
	}
}

func (repo *propertyRepository) GetPropertiesByProductId(ctx context.Context, productId int64) ([]coreEntity.Property, error) {
	var properties []gormModel.Property
	query := repo.gormDB.Model(&gormModel.Property{})
	if err := query.Where("product_id = ?", productId).Find(&properties).Error; err != nil {
		return nil, err
	}

	propertiesEntity := propertyRepo.MapListPropertiesGormModelToEntity(properties)
	return propertiesEntity, nil
}
func (repo *propertyRepository) CreateProperty(ctx context.Context, params domain.CreatePropertyRepoParams) (int64, error) {
	property := gormModel.Property{
		ProductId: params.ProductId,
		Name:      params.Name,
		Value:     params.Value,
	}
	if err := repo.gormDB.Create(&property).Error; err != nil {
		return 0, err
	}

	return property.Id, nil
}
func (repo *propertyRepository) UpdatePropertyById(ctx context.Context, params domain.UpdatePropertyByIdRepoParams) error {
	property := gormModel.Property{
		Name:  params.NewInfo.Name,
		Value: params.NewInfo.Value,
	}

	query := repo.gormDB.Model(&gormModel.Property{})
	if err := query.Updates(&property).Error; err != nil {
		return err
	}

	return nil

}
func (repo *propertyRepository) DeletePropertyById(ctx context.Context, propertyId int64) error {
	return repo.gormDB.Table(gormModel.Property{}.TableName()).Delete(&gormModel.Property{Id: propertyId}).Error
}
func (repo *propertyRepository) DeletePropertyByProductId(ctx context.Context, productId int64) error {
	return repo.gormDB.Table(gormModel.Property{}.TableName()).Where("product_id = ?", productId).Delete(&gormModel.Property{}).Error
}
