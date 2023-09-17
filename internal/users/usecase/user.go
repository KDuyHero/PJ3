package user_usecase

import (
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"
	coreModel "mobile-ecommerce/internal/core/model"
	"mobile-ecommerce/internal/util"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func newListUsersModel(users []coreEntity.User, pagination *common.Pagination) *coreModel.ListUsesModel {
	return &coreModel.ListUsesModel{
		Users:      mapListUsersEntityToModel(users),
		Pagination: *pagination,
	}
}

func (usecase *userUsecase) GetUsers(params domain.GetUsersUsecaseParams) (*coreModel.ListUsesModel, error) {
	// check role
	isAdmin := util.CheckRole(params.Role, domain.RequireAdmin)
	if !isAdmin {
		return nil, coreError.ErrNotPermission
	}

	//validate pagination
	common.CheckPagination(&params.Limit, &params.Page)
	repoParams := domain.GetUsersRepoParams{
		KeyWord: params.KeyWord,
		Limit:   params.Limit,
		Page:    params.Page,
		OrderBy: params.OrderBy,
	}
	// get users
	users, pagination, err := usecase.userRepo.GetUsers(repoParams)
	if err != nil {
		return nil, err
	}

	listUsersModel := newListUsersModel(users, pagination)
	return listUsersModel, nil
}

func (usecase *userUsecase) GetUserById(params domain.GetUserByIdUsecaseParams) (*coreModel.UserModel, error) {
	isAdmin := util.CheckRole(params.RoleUserRequest, domain.RequireAdmin)
	if params.IdUserRequest != params.IdUser && !isAdmin {
		return nil, coreError.ErrNotPermission
	}

	user, err := usecase.userRepo.GetUserById(params.Ctx, params.IdUser)
	if err != nil {
		return nil, err
	}
	userModel := mapUserEntityToModel(*user)
	return &userModel, nil
}

func (usecase *userUsecase) UpdateUserById(params domain.UpdateUserByIdUsecaseParams) error {
	// check role
	isAdmin := util.CheckRole(params.RoleUserRequest, domain.RequireAdmin)
	//if not self update and not admin
	if params.IdUserRequest != params.IdUserUpdate && !isAdmin {
		return coreError.ErrNotPermission
	}

	// update
	err := usecase.userRepo.UpdateUserById(domain.UpdateUserByIdRepoParams{
		Ctx:         params.Ctx,
		Id:          params.IdUserUpdate,
		Name:        params.Name,
		Avatar:      params.Avatar,
		PhoneNumber: params.PhoneNumber,
		UserName:    params.UserName,
	})

	if err != nil {
		return err
	}

	return nil
}
