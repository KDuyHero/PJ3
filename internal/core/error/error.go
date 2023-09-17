package coreError

import (
	"errors"
	"mobile-ecommerce/internal/common"
)

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrEmailExisted        = errors.New("email existed")
	ErrWrongPassword       = errors.New("wrong password")
	ErrNotPermission       = errors.New("don't have permission")
	ErrBrandNotFound       = errors.New("brand not found")
	ErrBrandNameExisted    = errors.New("brand'name existed")
	ErrBrandExisted        = errors.New("brand existed")
	ErrCategoryNotFound    = errors.New("category not found")
	ErrCategoryExisted     = errors.New("category existed")
	ErrCategoryNameExisted = errors.New("category'name existed")
	ErrCartNotFound        = errors.New("cart not found")
	ErrCartExisted         = errors.New("cart existed")
	ErrProductNotFound     = errors.New("product not found")
)

// user 1-50
// brand 51-100
// category 101-150
// cart 151-200
// product 201-250

func MapCoreErrToResponse(err error) common.BaseResponse {
	switch err {
	case ErrUserNotFound:
		return common.NewErrResponse(404_001, err.Error())
	case ErrInvalidEmail:
		return common.NewErrResponse(400_002, err.Error())
	case ErrEmailExisted:
		return common.NewErrResponse(400_003, err.Error())
	case ErrWrongPassword:
		return common.NewErrResponse(401_004, err.Error())
	case ErrNotPermission:
		return common.NewErrResponse(403_005, err.Error())
	case ErrBrandNotFound:
		return common.NewErrResponse(404_051, err.Error())
	case ErrBrandExisted:
		return common.NewErrResponse(400_052, err.Error())
	case ErrBrandNameExisted:
		return common.NewErrResponse(400_053, err.Error())
	case ErrCategoryNotFound:
		return common.NewErrResponse(404_101, err.Error())
	case ErrCategoryExisted:
		return common.NewErrResponse(400_102, err.Error())
	case ErrCategoryNameExisted:
		return common.NewErrResponse(400_103, err.Error())
	case ErrCategoryNotFound:
		return common.NewErrResponse(404_151, err.Error())
	case ErrCategoryExisted:
		return common.NewErrResponse(400_152, err.Error())
	case ErrProductNotFound:
		return common.NewErrResponse(404_201, err.Error())
	default:
		return common.NewErrResponse(500_000, err.Error())
	}
}

func GetHttpStatusFromCoreErrResponse(response common.BaseResponse) int {
	return response.Code / 1000
}
