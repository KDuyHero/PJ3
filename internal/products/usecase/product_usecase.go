package product_usecase

import (
	"context"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	coreModel "mobile-ecommerce/internal/core/model"
	"mobile-ecommerce/internal/util"
)

type productUsecase struct {
	productRepo  domain.ProductRepository
	commonRepo   domain.CommonRepository
	userRepo     domain.UserRepository
	propertyRepo domain.PropertyRepository
}

func NewProductUsecase(productRepo domain.ProductRepository, userRepo domain.UserRepository, propertyRepo domain.PropertyRepository, commonRepo domain.CommonRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepo:  productRepo,
		userRepo:     userRepo,
		propertyRepo: propertyRepo,
		commonRepo:   commonRepo,
	}
}

func (Uc *productUsecase) GetProducts(ctx context.Context, params domain.GetProductsUcParams) (*coreModel.ProductsModel, error) {
	common.CheckPagination(&params.Limit, &params.Page)

	products, pagination, err := Uc.productRepo.GetProducts(ctx, domain.GetProductsRepoParams{
		GetQueryFields: common.GetQueryFields{
			Keyword: params.Keyword,
			Limit:   params.Limit,
			Page:    params.Page,
		},
	})

	if err != nil {
		return nil, err
	}

	productsModel := MapListProductsEntityToCoreModel(*products)
	return &coreModel.ProductsModel{
		Products:   productsModel,
		Pagination: *pagination,
	}, nil
}

func (Uc *productUsecase) GetProductById(ctx context.Context, productId int64) (*coreModel.ProductModel, error) {
	product, err := Uc.productRepo.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}

	productModel := MapProductEntityToCoreModel(*product)
	return &productModel, nil
}

func (Uc *productUsecase) AddProduct(ctx context.Context, params domain.AddProductUcParams) error {
	// Check role
	user, _ := Uc.userRepo.GetUserById(ctx, params.UserId)
	if user == nil {
		return coreError.ErrUserNotFound
	}

	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	_, err := Uc.productRepo.CreateProduct(ctx, domain.CreateProductRepoParams{
		Name:       params.Name,
		Thumbnail:  params.Thumbnail,
		BrandName:  params.BrandName,
		Properties: params.Properties,
	})

	return err
}
func (Uc *productUsecase) UpdateProductById(ctx context.Context, params domain.UpdateProductParams) error {
	// Check role
	user, _ := Uc.userRepo.GetUserById(ctx, params.UserId)
	if user == nil {
		return coreError.ErrUserNotFound
	}

	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// check product existed
	product, _ := Uc.productRepo.GetProductById(ctx, params.ProductId)
	if product == nil {
		return coreError.ErrProductNotFound
	}

	// update product
	return Uc.productRepo.UpdateProductById(ctx, domain.UpdateProductByIdRepoParams{
		ProductId: params.ProductId,
		NewInfo: domain.UpdateProductInfo{
			Name:       params.Name,
			Thumbnail:  params.Thumbnail,
			BrandName:  params.BrandName,
			Properties: params.Properties,
		},
	})

}
func (Uc *productUsecase) DeleteProductById(ctx context.Context, userId int64, productId int64) error {
	// Check role
	user, _ := Uc.userRepo.GetUserById(ctx, userId)
	if user == nil {
		return coreError.ErrUserNotFound
	}

	isAdmin := util.CheckRole(user.Role, domain.RequireAdmin)
	if !isAdmin {
		return coreError.ErrNotPermission
	}

	// DeleteProduct
	return Uc.productRepo.DeleteProductById(ctx, productId)
}
