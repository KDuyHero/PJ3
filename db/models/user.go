package gormModel

import (
	"database/sql"
	"time"
)

type User struct {
	Id                int64          `gorm:"primary_key;column:id;"`
	Name              string         `gorm:"column:name"`
	Username          string         `gorm:"column:username"`
	EncryptedPassword string         `gorm:"column:encrypted_password"`
	Email             string         `gorm:"column:email"`
	PhoneNumber       sql.NullString `gorm:"column:phone_number"`
	Avatar            sql.NullString `gorm:"column:avatar"`
	Role              int8           `gorm:"column:role"`
	Status            string         `gorm:"column:status"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
}

func (User) TableName() string { return "users" }

type UserCreate struct {
	Id                int64  `gorm:"primary_key;column:id;"`
	Name              string `gorm:"column:name"`
	Username          string `gorm:"column:username"`
	EncryptedPassword string `gorm:"column:encrypted_password"`
	Email             string `gorm:"column:email"`
}

func (UserCreate) TableName() string { return User{}.TableName() }
