package util

import (
	"mobile-ecommerce/internal/core/domain"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func GetUserInfoFromContext(ctx *gin.Context) (*token.Payload, bool) {
	value, ok := ctx.Get(middleware.AuthorizationPayloadKey)
	if !ok {
		return nil, false
	}

	payload, ok := value.(*token.Payload)
	return payload, ok
}

func CheckRole(userRole int8, requireRoles domain.RequireRoles) bool {
	role := MapNumberToRoleUserType(userRole)
	for _, value := range requireRoles {
		if role == value {
			return true
		}
	}
	return false
}

func MapNumberToRoleUserType(number int8) domain.RoleUser {
	switch number {
	case 1:
		return domain.CustomerRole
	case 2:
		return domain.AdminRole
	default:
		return domain.GuestRole
	}
}
