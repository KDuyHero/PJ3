package cartRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapCartGormModelToEntity(cart gormModel.Cart) coreEntity.Cart {
	return coreEntity.Cart{
		Id:     cart.Id,
		UserId: cart.UserId,
	}
}
