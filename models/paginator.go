package models

import "math"

type Pagination struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalPage int `json:"total_page"`
	TotalDocs int `json:"total_docs"`
}

func Paginator(paginator *Pagination) *Pagination {
	return &Pagination{
		Page:      paginator.Page,
		PerPage:   paginator.PerPage,
		TotalPage: int(math.Ceil(float64(paginator.TotalDocs) / float64(paginator.PerPage))),
		TotalDocs: paginator.TotalDocs,
	}
}
