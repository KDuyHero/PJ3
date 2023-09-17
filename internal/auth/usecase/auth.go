package auth_usecase

import (
	"fmt"
	"mobile-ecommerce/internal/core/domain"
	coreError "mobile-ecommerce/internal/core/error"
	"mobile-ecommerce/internal/util"
	"mobile-ecommerce/internal/util/token"
	"net/mail"
	"strconv"

	"github.com/gosimple/slug"
)

type authUsecase struct {
	userRepo domain.UserRepository
	jwtMaker token.TokenMaker
}

func NewAuthUsecase(userRepo domain.UserRepository, jwtMaker token.TokenMaker) domain.AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		jwtMaker: jwtMaker,
	}

}

func (authUsecase *authUsecase) Login(params domain.LoginParams) (string, *domain.UserInfoResponse, error) {
	// validate email
	if _, err := mail.ParseAddress(params.Email); err != nil {
		return "", nil, coreError.ErrInvalidEmail
	}

	// check email existed
	user, err := authUsecase.userRepo.GetUserByEmail(params.Ctx, params.Email)
	if err != nil {
		return "", nil, err
	}
	// compare password
	if err := util.CheckPassword(user.EncryptedPassword, params.Password); err != nil {
		return "", nil, err
	}
	// gen token
	tokenString, err := authUsecase.jwtMaker.GenerateToken(user.Id, user.Name, user.Role)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &domain.UserInfoResponse{
		Name: user.Name,
		Role: user.Role,
	}, nil
}

func (authUsecase *authUsecase) Register(params domain.RegisterParams) error {
	// validate email
	_, err := mail.ParseAddress(params.Email)
	if err != nil {
		return coreError.ErrInvalidEmail
	}

	// check email has been exists
	user, _ := authUsecase.userRepo.GetUserByEmail(params.Ctx, params.Email)
	if user != nil {
		return coreError.ErrEmailExisted
	}

	// validate password (strong or week)
	// encrypt password
	encryptedPassword, err := util.EncryptPassword(params.Password)
	if err != nil {
		return err
	}
	// gen username
	username := slug.Make(params.Name)
	user, _ = authUsecase.userRepo.GetUserByUsername(params.Ctx, username)
	if user != nil {
		x := 1
		for {
			user, _ = authUsecase.userRepo.GetUserByUsername(params.Ctx, username+"-"+strconv.Itoa(x))
			if user == nil {
				username += "-" + strconv.Itoa(x)
				break
			}
			x += 1
		}
	}

	_, err = authUsecase.userRepo.CreateUser(domain.CreateUserRepoParams{
		Ctx:               params.Ctx,
		Name:              params.Name,
		Username:          username,
		Email:             params.Email,
		EncryptedPassword: encryptedPassword,
	})

	if err != nil {
		return err
	}
	fmt.Println(4)

	return nil
}
