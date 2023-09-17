package cart_usecase

import (
	"context"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	coreModel "mobile-ecommerce/internal/core/model"
)

type cartUsecase struct {
	cartRepo domain.CartRepository
	userRepo domain.UserRepository
}

func NewCartUsecase(cartRepo domain.CartRepository, userRepo domain.UserRepository) domain.CartUsecase {
	return &cartUsecase{
		cartRepo: cartRepo,
		userRepo: userRepo,
	}
}

func (Uc *cartUsecase) GetCart(ctx context.Context, userId int64) (*coreModel.CartModel, error) {
	// check user existed
	user, _ := Uc.userRepo.GetUserById(ctx, userId)
	if user == nil {
		return nil, coreError.ErrUserNotFound
	}

	// get cart
	cart, err := Uc.cartRepo.GetCartByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	cartModel := mapCartEntityToCoreModel(*cart)
	return &cartModel, nil

}
func (Uc *cartUsecase) CreateCart(ctx context.Context, params domain.AddCartUcParams) error {
	// check user existed
	user, _ := Uc.userRepo.GetUserById(ctx, params.UserId)
	if user == nil {
		return coreError.ErrUserNotFound
	}

	// create cart
	_, err := Uc.cartRepo.CreateCart(ctx, domain.AddCartRepoParams{
		UserId: user.Id,
	})
	if err != nil {
		return err
	}

	return nil
}
func (Uc *cartUsecase) DeleteCart(ctx context.Context, userId int64) error {
	// check user existed
	return Uc.cartRepo.DeleteCartByUserId(ctx, userId)
}
