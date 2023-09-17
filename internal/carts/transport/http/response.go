package cart_http

import coreModel "mobile-ecommerce/internal/core/model"

type getCartResponse struct {
	Code int
	Cart coreModel.CartModel
}

func NewGetCartResponse(cart coreModel.CartModel) getCartResponse {
	return getCartResponse{
		Code: 0,
		Cart: cart,
	}
}
