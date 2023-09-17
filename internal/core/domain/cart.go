package domain

import (
	"context"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

type AddCartRepoParams struct {
	UserId int64
}

type CartRepository interface {
	GetCartById(ctx context.Context, cartId int64) (*coreEntity.Cart, error)
	GetCartByUserId(ctx context.Context, userId int64) (*coreEntity.Cart, error)
	CreateCart(ctx context.Context, params AddCartRepoParams) (int64, error)
	DeleteCartById(ctx context.Context, cartId int64) error
	DeleteCartByUserId(ctx context.Context, userId int64) error
}

type AddCartUcParams struct {
	UserId int64
}
type CartUsecase interface {
	GetCart(ctx context.Context, userId int64) (*coreModel.CartModel, error)
	CreateCart(ctx context.Context, params AddCartUcParams) error
	DeleteCart(ctx context.Context, userId int64) error
}
