package domain

import "context"

type LoginParams struct {
	Ctx      context.Context
	Email    string
	Password string
}
type RegisterParams struct {
	Ctx      context.Context
	Name     string
	Email    string
	Password string
}

type UserInfoResponse struct {
	Name string
	Role int8
}
type AuthUsecase interface {
	Login(params LoginParams) (string, *UserInfoResponse, error)
	Register(params RegisterParams) error
}
