package meta

import (
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
}

func New(page, perPage, total int, pageLimDef string) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pageLimDef)
		if err != nil {
			return nil, err
		}
	}

	//Obtenemos el total de paginas
	pageCount := 0
	if total > 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	// Validamos que no nos pasen un page menor a 1
	if page < 1 {
		page = 1
	}
	return &Meta{TotalCount: total,
		Page:      page,
		PerPage:   perPage,
		PageCount: pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}
