package domain

import "context"

type CartProductCategory interface {
	AddProductToCart(ctx context.Context, cartId int64, productId int64) error
	UpdateQuantity(ctx context.Context, cartId int64, productId int64, quantity int32) error
	RemoveProduct(ctx context.Context, cartId int64, productId int64) error
	ClearCart(ctx context.Context, cartId int64) error
}
type CartProductUsecase interface {
	AddProductToCart(ctx context.Context, cartId int64, productId int64) error
	UpdateQuantity(ctx context.Context, cartId int64, productId int64, quantity int32) error
	RemoveProduct(ctx context.Context, cartId int64, productId int64) error
	ClearCart(ctx context.Context, cartId int64) error
}
