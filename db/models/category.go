package gormModel

import (
	"time"
)

type Category struct {
	Name      string    `gorm:"primary_key;column:name"`
	Slug      string    `gorm:"column:slug"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
