package gormModel

import "time"

type CartProduct struct {
	CartId    int64     `gorm:"primary_key;column:cart_id"`
	ProductId int64     `gorm:"primary_key;column:product_id"`
	Quantity  int32     `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (CartProduct) TableName() string { return "carts_products" }
