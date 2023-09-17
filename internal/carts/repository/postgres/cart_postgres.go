package cart_postgres

import (
	"context"
	gormModel "mobile-ecommerce/db/models"
	cartRepo "mobile-ecommerce/internal/carts/repository"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"

	"gorm.io/gorm"
)

type cartRepository struct {
	gormDB *gorm.DB
}

func NewCartRepository(gormDB *gorm.DB) domain.CartRepository {
	return &cartRepository{
		gormDB: gormDB,
	}
}

func (repo *cartRepository) GetCartById(ctx context.Context, cartId int64) (*coreEntity.Cart, error) {
	cart := gormModel.Cart{}
	if err := repo.gormDB.Where("id = ?", cartId).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrCartNotFound
		}
		return nil, err
	}

	cartEntity := cartRepo.MapCartGormModelToEntity(cart)
	return &cartEntity, nil
}
func (repo *cartRepository) GetCartByUserId(ctx context.Context, userId int64) (*coreEntity.Cart, error) {
	cart := gormModel.Cart{}
	if err := repo.gormDB.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrCartNotFound
		}
		return nil, err
	}

	cartEntity := cartRepo.MapCartGormModelToEntity(cart)
	return &cartEntity, nil
}
func (repo *cartRepository) CreateCart(ctx context.Context, params domain.AddCartRepoParams) (int64, error) {
	newCart := gormModel.Cart{
		UserId: params.UserId,
	}
	if err := repo.gormDB.Create(&newCart).Error; err != nil {
		return 0, err
	}

	return newCart.Id, nil
}
func (repo *cartRepository) DeleteCartById(ctx context.Context, cartId int64) error {
	return repo.gormDB.Table(gormModel.Cart{}.TableName()).Where("id = ?", cartId).Delete(&gormModel.Cart{}).Error
}
func (repo *cartRepository) DeleteCartByUserId(ctx context.Context, userId int64) error {
	return repo.gormDB.Table(gormModel.Cart{}.TableName()).Where("user_id = ?", userId).Delete(&gormModel.Cart{}).Error

}
