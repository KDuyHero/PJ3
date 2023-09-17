package cart_usecase

import (
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

func mapCartEntityToCoreModel(cart coreEntity.Cart) coreModel.CartModel {
	return coreModel.CartModel{
		Id:     cart.Id,
		UserId: cart.UserId,
	}
}
