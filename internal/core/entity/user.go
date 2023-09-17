package coreEntity

import (
	"time"
)

type User struct {
	Id                int64
	Name              string
	Username          string
	EncryptedPassword string
	Email             string
	PhoneNumber       string
	Avatar            string
	Role              int8
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
