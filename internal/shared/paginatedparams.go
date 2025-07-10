package shared

type PaginatedQueryParamsDTO struct {
	PerPage        int64  `json:"perPage" validate:"gte=0"`
	Page           int64  `json:"page" validate:"gte=0"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection" default:"DESC" validate:"oneof=ASC DESC"`
	SearchFilter   string `json:"searchFilter"`
}
type PaginatedQueryParams struct {
	Limit          int64
	Offset         int64
	OrderBy        string
	OrderDirection string
	SearchFilter   string
}

func (p *PaginatedQueryParamsDTO) DefaultValues() {
	if p.PerPage == 0 {
		p.PerPage = 10
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.OrderBy == "" {
		p.OrderBy = "created_at"
	}
	if p.OrderDirection == "" {
		p.OrderDirection = "DESC"
	}
}

func (p *PaginatedQueryParamsDTO) DTOToModel() *PaginatedQueryParams {
	p.DefaultValues()
	return &PaginatedQueryParams{
		Limit:          p.PerPage,
		Offset:         (p.Page - 1) * p.PerPage,
		OrderBy:        p.OrderBy,
		OrderDirection: p.OrderDirection,
		SearchFilter:   p.SearchFilter,
	}

}
