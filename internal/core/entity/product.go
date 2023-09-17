package coreEntity

type Product struct {
	Id         int64
	Name       string
	Thumbnail  string
	BrandName  string
	Properties []struct {
		Name  string
		Value string
	}
}
