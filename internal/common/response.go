package common

import "net/http"

type BaseResponse struct {
	Code    int
	Message string
}

func NewErrResponse(code int, message string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
	}
}

func NewSuccessResponse() BaseResponse {
	return BaseResponse{
		Code:    0,
		Message: "success",
	}
}

func NewUnAuthorizeResponse() BaseResponse {
	return BaseResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	}
}
