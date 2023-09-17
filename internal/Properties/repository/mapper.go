package propertyRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapPropertyGormModelToEntity(property gormModel.Property) coreEntity.Property {
	return coreEntity.Property{
		Id:        property.Id,
		ProductId: property.ProductId,
		Name:      property.Name,
		Value:     property.Value,
	}
}

func MapListPropertiesGormModelToEntity(listProperties []gormModel.Property) []coreEntity.Property {
	var properties []coreEntity.Property

	for _, property := range listProperties {
		properties = append(properties, MapPropertyGormModelToEntity(property))
	}

	return properties
}
