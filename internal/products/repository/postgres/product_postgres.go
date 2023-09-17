package product_postgres

import (
	"context"
	"database/sql"
	"fmt"
	gormModel "mobile-ecommerce/db/models"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"
	productRepo "mobile-ecommerce/internal/products/repository"

	"gorm.io/gorm"
)

type productRepository struct {
	gormDB *gorm.DB
}

// get instance of productRepo
func NewProductRepository(gormDB *gorm.DB) domain.ProductRepository {
	return &productRepository{
		gormDB: gormDB,
	}
}

// get list product with condition
func (repo *productRepository) GetProducts(ctx context.Context, params domain.GetProductsRepoParams) (*[]coreEntity.Product, *common.Pagination, error) {
	var products []gormModel.Product
	query := repo.gormDB.Model(&gormModel.Product{})
	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		query.Where("name ILIKE ?", keyword)
		query.Where("brand ILIKE ?", keyword)
	}

	var totalRows int64
	if err := query.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	query.Preload("Properties")
	offset := (params.Page - 1) * common.LimitPerPageDefault
	if err := query.Offset(int(offset)).Limit(int(params.Limit)).Find(&products).Error; err != nil {
		return nil, nil, err
	}

	pagination := common.Pagination{
		Limit:     params.Limit,
		Page:      params.Page,
		TotalRows: totalRows,
	}
	productsEntity := productRepo.MapListProductsGormToEntity(products)

	return &productsEntity, &pagination, nil
}

// get a product with specific product's id
func (repo *productRepository) GetProductById(ctx context.Context, productId int64) (*coreEntity.Product, error) {
	var product gormModel.Product

	query := repo.gormDB.Model(&gormModel.Product{})
	query.Preload("Properties")
	if err := query.Where("id = ?", productId).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrProductNotFound
		}
		return nil, err
	}

	fmt.Println("product:", product)

	productEntity := productRepo.MapProductGormModelToEntity(product)
	return &productEntity, nil
}

// create new product
func (repo *productRepository) CreateProduct(ctx context.Context, params domain.CreateProductRepoParams) (int64, error) {
	product := gormModel.Product{
		Name: params.Name,
		Thumbnail: sql.NullString{
			String: params.Thumbnail,
			Valid:  len(params.Thumbnail) > 0,
		},
		BrandName: params.BrandName,
	}

	err := repo.gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}
		for _, property := range params.Properties {
			if err := tx.Create(&gormModel.Property{
				ProductId: product.Id,
				Name:      property.Name,
				Value:     property.Value,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return product.Id, nil
}

// update product with product's id
func (repo *productRepository) UpdateProductById(ctx context.Context, params domain.UpdateProductByIdRepoParams) error {
	return repo.gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&gormModel.Product{}).Where("id = ?", params.ProductId).Updates(&gormModel.Product{
			Name: params.NewInfo.Name,
			Thumbnail: sql.NullString{
				String: params.NewInfo.Thumbnail,
				Valid:  len(params.NewInfo.Thumbnail) > 0,
			},
			BrandName: params.NewInfo.BrandName,
		}).Error; err != nil {
			return err
		}

		// delete old properties
		if err := tx.Table(gormModel.Property{}.TableName()).Where("product_id = ?", params.ProductId).Delete(&gormModel.Property{}).Error; err != nil {
			return err
		}

		// create new properties
		for _, property := range params.NewInfo.Properties {
			if err := tx.Create(&gormModel.Property{
				ProductId: params.ProductId,
				Name:      property.Name,
				Value:     property.Value,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// delete product with product's id
func (repo *productRepository) DeleteProductById(ctx context.Context, productId int64) error {
	return repo.gormDB.Table(gormModel.Product{}.TableName()).Where("id = ?", productId).Delete(&gormModel.Product{}).Error
}
