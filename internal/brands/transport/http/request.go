package brand_http

type getBrandsRequest struct {
	KeyWord string `form:"q"`
	Limit   int32  `form:"limit"`
	Page    int32  `form:"page"`
}

type addNewBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type updateBrandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
