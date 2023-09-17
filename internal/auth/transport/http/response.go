package auth_http

import "mobile-ecommerce/internal/core/domain"

type loginSuccessResponse struct {
	Code  int
	Token string
	User  domain.UserInfoResponse
}

func newLoginSuccessResponse(token string, info domain.UserInfoResponse) loginSuccessResponse {
	return loginSuccessResponse{
		Code:  0,
		Token: token,
		User:  info,
	}
}
