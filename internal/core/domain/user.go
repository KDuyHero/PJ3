package domain

import (
	"context"
	"mobile-ecommerce/internal/common"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreModel "mobile-ecommerce/internal/core/model"
)

type RoleUser int8
type RequireRoles []RoleUser

const (
	CustomerRole RoleUser = 1
	AdminRole    RoleUser = 2
	GuestRole    RoleUser = 0
)

var (
	RequireCustomer RequireRoles = []RoleUser{
		CustomerRole,
		AdminRole,
	}
	RequireAdmin RequireRoles = []RoleUser{AdminRole}
)

type GetUsersRepoParams struct {
	Ctx     context.Context
	KeyWord string
	Limit   int32
	Page    int32
	OrderBy string
}

type CreateUserRepoParams struct {
	Ctx               context.Context
	Name              string
	Username          string
	Email             string
	EncryptedPassword string
}

type UpdateUserByIdRepoParams struct {
	Ctx         context.Context
	Id          int64
	Name        string
	UserName    string
	PhoneNumber string
	Avatar      string
}
type UserRepository interface {
	GetUsers(params GetUsersRepoParams) ([]coreEntity.User, *common.Pagination, error)
	GetUserById(ctx context.Context, id int64) (*coreEntity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*coreEntity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*coreEntity.User, error)
	CreateUser(params CreateUserRepoParams) (int64, error)
	UpdateUserById(params UpdateUserByIdRepoParams) error
}

type GetUsersUsecaseParams struct {
	Ctx     context.Context
	KeyWord string
	Limit   int32
	Page    int32
	OrderBy string
	Role    int8
}

type GetUserByIdUsecaseParams struct {
	Ctx             context.Context
	IdUser          int64
	IdUserRequest   int64
	RoleUserRequest int8
}

type UpdateUserByIdUsecaseParams struct {
	Ctx             context.Context
	IdUserRequest   int64
	IdUserUpdate    int64
	RoleUserRequest int8
	Name            string
	UserName        string
	PhoneNumber     string
	Avatar          string
}
type UserUsecase interface {
	GetUsers(params GetUsersUsecaseParams) (*coreModel.ListUsesModel, error)
	UpdateUserById(params UpdateUserByIdUsecaseParams) error
	GetUserById(params GetUserByIdUsecaseParams) (*coreModel.UserModel, error)
}
