package domain

import (
	"context"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

type CreatePropertyRepoParams struct {
	ProductId int64
	Name      string
	Value     string
}

type UpdatePropertyInfo struct {
	Name  string
	Value string
}
type UpdatePropertyByIdRepoParams struct {
	PropertyId int64
	NewInfo    UpdatePropertyInfo
}
type PropertyRepository interface {
	GetPropertiesByProductId(ctx context.Context, productId int64) ([]coreEntity.Property, error)
	CreateProperty(ctx context.Context, params CreatePropertyRepoParams) (int64, error)
	UpdatePropertyById(ctx context.Context, params UpdatePropertyByIdRepoParams) error
	DeletePropertyById(ctx context.Context, propertyId int64) error
	DeletePropertyByProductId(ctx context.Context, productId int64) error
}
