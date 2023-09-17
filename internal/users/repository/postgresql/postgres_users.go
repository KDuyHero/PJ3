package user_postgres

import (
	"context"
	"database/sql"
	gormModel "mobile-ecommerce/db/models"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/core/domain"
	coreEntity "mobile-ecommerce/internal/core/entity"
	coreError "mobile-ecommerce/internal/core/error"
	userRepo "mobile-ecommerce/internal/users/repository"
	"mobile-ecommerce/internal/util"

	"gorm.io/gorm"
)

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		gormDB: db,
	}
}

func (repo *userRepository) GetUsers(params domain.GetUsersRepoParams) ([]coreEntity.User, *common.Pagination, error) {
	var listUsersGorm []gormModel.User
	query := repo.gormDB.Model(&gormModel.User{})
	if params.KeyWord != "" {
		keyWord := "%" + params.KeyWord + "%"
		query.Where("name ilike ?", keyWord)
		query.Where("username ilike ?", keyWord)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	offset := (params.Page - 1) * util.DEFAULT_PER_PAGE
	if params.OrderBy != "" {
		query.Order(params.OrderBy)
	}

	if err := query.Offset(int(offset)).Limit(int(params.Limit)).Find(&listUsersGorm).Error; err != nil {
		return nil, nil, err
	}
	pagination := &common.Pagination{
		Limit:     int32(params.Limit),
		Page:      int32(params.Page),
		TotalRows: total,
	}

	return userRepo.MapListUsersGormToEntity(listUsersGorm), pagination, nil

}

func (repo *userRepository) GetUserById(ctx context.Context, id int64) (*coreEntity.User, error) {
	var user gormModel.User
	query := repo.gormDB.Model(&gormModel.User{})

	if err := query.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrUserNotFound
		}
		return nil, err
	}

	return userRepo.MapUserGormToEntity(user), nil
}

func (repo *userRepository) GetUserByEmail(ctx context.Context, email string) (*coreEntity.User, error) {
	var user gormModel.User
	query := repo.gormDB.Model(&gormModel.User{})
	if err := query.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrUserNotFound
		}
		return nil, err
	}

	return userRepo.MapUserGormToEntity(user), nil
}

func (repo *userRepository) GetUserByUsername(ctx context.Context, username string) (*coreEntity.User, error) {
	var user gormModel.User
	query := repo.gormDB.Model(&gormModel.User{})
	if err := query.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, coreError.ErrUserNotFound
		}
		return nil, err
	}

	return userRepo.MapUserGormToEntity(user), nil
}

func (repo *userRepository) CreateUser(params domain.CreateUserRepoParams) (int64, error) {
	user := gormModel.UserCreate{
		Name:              params.Name,
		Username:          params.Username,
		EncryptedPassword: params.EncryptedPassword,
		Email:             params.Email,
	}

	if err := repo.gormDB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (repo *userRepository) UpdateUserById(params domain.UpdateUserByIdRepoParams) error {
	newInfo := gormModel.User{
		Name:     params.Name,
		Username: params.UserName,
		PhoneNumber: sql.NullString{
			String: params.PhoneNumber,
			Valid:  len(params.PhoneNumber) > 0,
		},
		Avatar: sql.NullString{
			String: params.Avatar,
			Valid:  len(params.Avatar) > 0,
		},
	}
	return repo.gormDB.Model(&gormModel.User{Id: params.Id}).Updates(&newInfo).Error
}
