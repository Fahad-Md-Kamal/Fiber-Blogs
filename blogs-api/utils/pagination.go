package utils

import (
	"math"
)

type Pagination struct {
	TotalCount  int64       `json:"total_count"`
	Limit       int         `json:"limit"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	HasNextPage bool        `json:"has_next_page"`
	HasPrevPage bool        `json:"has_prev_page"`
	NextPage    int         `json:"next_page"`
	PrevPage    int         `json:"prev_page"`
	Data        interface{} `json:"data"`
}

func Paginate(totalCount, limit, currentPage int, data interface{}) *Pagination {
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	hasNextPage := currentPage < totalPages
	hasPrevPage := currentPage > 1
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	return &Pagination{
		TotalCount:  int64(totalCount),
		Limit:       limit,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		Data:        data,
	}
}
