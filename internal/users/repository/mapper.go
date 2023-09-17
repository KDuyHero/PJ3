package userRepo

import (
	gormModel "mobile-ecommerce/db/models"
	coreEntity "mobile-ecommerce/internal/core/entity"
)

func MapUserGormToEntity(user gormModel.User) *coreEntity.User {
	return &coreEntity.User{
		Id:                user.Id,
		Name:              user.Name,
		Username:          user.Username,
		EncryptedPassword: user.EncryptedPassword,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber.String,
		Avatar:            user.Avatar.String,
		Role:              user.Role,
		Status:            user.Status,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}

func MapListUsersGormToEntity(users []gormModel.User) []coreEntity.User {
	var listUsersEntity []coreEntity.User
	for _, user := range users {
		listUsersEntity = append(listUsersEntity, *MapUserGormToEntity(user))
	}

	return listUsersEntity
}
