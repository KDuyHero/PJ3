package gormModel

import (
	"database/sql"
	"time"
)

type Brand struct {
	Name        string         `gorm:"primary_key;column:name"`
	Slug        string         `gorm:"column:slug"`
	Description sql.NullString `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
}

func (Brand) TableName() string { return "brands" }
