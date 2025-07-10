package query

import (
	"net/url"
	"strconv"
)

const (
	DefaultPerPage = 20
	MaxPerPage     = 1000
)

const (
	PageQueryParam    = "page"
	PerPageQueryParam = "per_page"
	SortQueryParam    = "sort"
	FilterQueryParam  = "filter"
	ExpandQueryParam  = "expand"
)

type QueryDTO struct {
	Page    int
	PerPage int
	Filter  string
	Expand  string
	Sort    string
}

func NewQueryDTO() *QueryDTO {
	return &QueryDTO{
		Page:    1,
		PerPage: DefaultPerPage,
		Filter:  "",
		Expand:  "",
		Sort:    "",
	}
}

func (q *QueryDTO) Parse(urlQuery string) error {
	params, err := url.ParseQuery(urlQuery)
	if err != nil {
		return err
	}

	if raw := params.Get(PageQueryParam); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return err
		}
		q.Page = v
	}

	if raw := params.Get(PerPageQueryParam); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return err
		}
		q.PerPage = v
	}
	if raw := params.Get(FilterQueryParam); raw != "" {
		q.Filter = raw
	}
	if raw := params.Get(ExpandQueryParam); raw != "" {
		q.Expand = raw
	}
	if raw := params.Get(SortQueryParam); raw != "" {
		q.Sort = raw
	}
	return nil
}
