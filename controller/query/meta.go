package query

type Pagination struct {
	Page     int64 `json:"page" schema:"page"`
	PageSize int64 `json:"page_size" schema:"page_size"`
}
