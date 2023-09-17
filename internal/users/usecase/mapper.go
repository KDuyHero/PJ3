package user_usecase

import (
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

func mapUserEntityToModel(user coreEntity.User) coreModel.UserModel {
	return coreModel.UserModel{
		Id:          user.Id,
		Name:        user.Name,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Status:      user.Status,
	}
}

func mapListUsersEntityToModel(users []coreEntity.User) []coreModel.UserModel {
	var usersModel []coreModel.UserModel
	for _, user := range users {
		usersModel = append(usersModel, mapUserEntityToModel(user))
	}

	return usersModel
}
