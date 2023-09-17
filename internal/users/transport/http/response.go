package user_http

import coreModel "mobile-ecommerce/internal/core/model"

type getUsersResponse struct {
	Code int
	Data coreModel.ListUsesModel
}

type getUserByIdResponse struct {
	Code int
	Data coreModel.UserModel
}

func newGetUsersResponse(model coreModel.ListUsesModel) getUsersResponse {
	return getUsersResponse{
		Code: 0,
		Data: model,
	}
}

func NewGetUserByIdResponse(model coreModel.UserModel) getUserByIdResponse {
	return getUserByIdResponse{
		Code: 0,
		Data: model,
	}
}
