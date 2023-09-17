package gormModel

import "time"

type Cart struct {
	Id        int64     `gorm:"primary_key;column:id"`
	UserId    int64     `gorm:"colum:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Owner User `gorm:"foreignKey:UserId"`
}

func (Cart) TableName() string { return "carts" }
