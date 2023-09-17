package middleware

import (
	"errors"
	"mobile-ecommerce/internal/common"
	"mobile-ecommerce/internal/util/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	AuthorizationHeader     = "Authorization"
	AuthorazationType       = "Bearer"
	AuthorizationPayloadKey = "user-info"

	errTokenMissing = errors.New("token missing")
	errInvalidToken = errors.New("invalid token")
)

func RequireAuthMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := verifyAuthorizationHeader(ctx, tokenMaker, true)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, common.NewErrResponse(http.StatusForbidden, err.Error()))
			return
		}
		ctx.Next()
	}
}

func OptionalAuthMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyAuthorizationHeader(ctx, tokenMaker, false)
		ctx.Next()
	}
}
func verifyAuthorizationHeader(ctx *gin.Context, tokenMaker token.TokenMaker, isRequire bool) error {
	authHeader := ctx.GetHeader(AuthorizationHeader)
	if len(authHeader) == 0 {
		return errTokenMissing
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 {
		return errInvalidToken
	}

	authorizationType := fields[0]
	if authorizationType != AuthorazationType {
		return errInvalidToken
	}

	authToken := fields[1]
	payload, err := tokenMaker.VerifyToken(authToken)
	if err != nil {
		if isRequire {
			return errInvalidToken
		} else {
			ctx.Set(AuthorizationPayloadKey, token.NewPayLoad(0, "", 0, 0))
		}
	} else {
		ctx.Set(AuthorizationPayloadKey, payload)
	}

	return nil

}
