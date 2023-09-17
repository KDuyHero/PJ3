package user_http

type GetUsersRequest struct {
	KeyWord string `form:"q"`
	Limit   int32  `form:"limit"`
	Page    int32  `form:"page"`
	OrderBy string `form:"orderBy"`
}
