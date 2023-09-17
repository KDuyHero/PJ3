package common

const (
	LimitPerPageDefault = 10
	PageDefaut          = 1
)

type Pagination struct {
	Limit     int32 `json:"limit"`
	Page      int32 `json:"page"`
	TotalRows int64 `json:"totalRows"`
}

func CheckPagination(limit *int32, page *int32) {
	if *limit <= 0 {
		*limit = LimitPerPageDefault
	}

	if *page <= 0 {
		*page = PageDefaut
	}
}
