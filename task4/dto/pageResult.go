package dto

type PageResult[T any] struct {
	Data      []T   `json:"data"`
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalPage int   `json:"total_page"`
}
