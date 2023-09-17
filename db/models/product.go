package gormModel

import (
	"database/sql"
	"time"
)

type Product struct {
	Id        int64          `gorm:"primary_key;column:id"`
	Name      string         `gorm:"column:name"`
	Thumbnail sql.NullString `gorm:"column:thumbnail"`
	BrandName string         `gorm:"column:brand"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`

	Brand      Brand      `gorm:"foreignKey:BrandName"`
	Properties []Property `gorm:"foreignKey:ProductId;references:Id"`
}

func (Product) TableName() string { return "products" }
