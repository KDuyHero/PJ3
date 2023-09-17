package coreModel

import (
	"mobile-ecommerce/internal/common"
)

type UserModel struct {
	Id          int64
	Name        string
	Username    string
	Email       string
	PhoneNumber string
	Avatar      string
	Role        int8
	Status      string
}

type ListUsesModel struct {
	Users      []UserModel
	Pagination common.Pagination
}
