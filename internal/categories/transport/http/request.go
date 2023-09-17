package category_http

type getCategoriesRequest struct {
	Keyword string `form:"q"`
	Limit   int32  `form:"limit"`
	Page    int32  `form:"page"`
}

type addCategoryRequest struct {
	Name string `json:"name"`
}

type updateCategoryRequest struct {
	Name string `json:"name"`
}
