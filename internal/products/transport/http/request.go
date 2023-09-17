package product_http

type getProductsRequest struct {
	Keyword string `form:"q"`
	Limit   int32  `form:"limit"`
	Page    int32  `form:"page"`
}
type addProductRequest struct {
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	BrandName  string `json:"brand"`
	Properties []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
}

type updateProductRequest struct {
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	BrandName  string `json:"brand"`
	Properties []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
}
