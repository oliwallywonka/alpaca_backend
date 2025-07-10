package tourfx

import (
	"database/sql"
	/* "errors" */
	"fmt"

	/* "github.com/go-jet/jet/v2/qrm" */
	. "github.com/go-jet/jet/v2/sqlite"
	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type IService interface {
	GetAll(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*model.Tour, error)
	Save(dto *CreateTourDTO) (*model.Tour, error)
	Update(tourID string, dto *model.Tour) (*model.Tour, error)
	NameOrSlugExists(name string) (bool, error)

	
}

type service struct {
	db   *sql.DB
	gorm *gorm.DB
}

// GetAll implements IService.
func (s *service) GetAll(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	var tours []model.Tour = make([]model.Tour, 0)
	var countStruct []struct {
		Total int
	}
	params := dto.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := Tour.SELECT(Tour.AllColumns).
		WHERE(
			OR(
				Tour.Code.LIKE(String(searchFilter)),
				CAST(Tour.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Tour.Slug).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Tour.Days).AS_TEXT().LIKE(String(searchFilter)),
				Tour.GroupSize.LIKE(String(searchFilter)),
			),
		).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &tours)
	if err != nil {
		return nil, err
	}
	stmCount := Tour.SELECT(COUNT(STAR).AS("total")).
		WHERE(
			OR(
				Tour.Code.LIKE(String(searchFilter)),
				CAST(Tour.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Tour.Slug).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Tour.Days).AS_TEXT().LIKE(String(searchFilter)),
				Tour.GroupSize.LIKE(String(searchFilter)),
			),
		)
	err = stmCount.Query(s.db, &countStruct)
	if err != nil {
		return nil, err
	}
	var totalUsers int = 0
	if len(countStruct) > 0 {
		totalUsers = countStruct[0].Total
	}
	return &shared.PaginatedResponse{
		Items: tours,
		Page:  int(dto.Page),
		Total: totalUsers,
	}, nil
}

// GetByUniqueKey implements IService.
func (s *service) GetByUniqueKey(key string) (*model.Tour, error) {
	panic("unimplemented")
}

// NameOrSlugExists implements IService.
func (s *service) NameOrSlugExists(name string) (bool, error) {
	panic("unimplemented")
}

// Save implements IService.
func (s *service) Save(dto *CreateTourDTO) (*model.Tour, error) {
	tour := dto.DTOToModel()
	stm := Tour.INSERT(
		Tour.ID,
		Tour.Code,
		Tour.Name,
		Tour.Slug,
		Tour.Transport,
		Tour.Accommodation,
		Tour.Team,
		Tour.ShortDescription,
		Tour.LongDescription,
		Tour.Images,
		Tour.CreatedAt,
		Tour.UpdatedAt,
	).
		MODEL(tour).
		RETURNING(Tour.AllColumns)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, err
	}
	return tour, nil
}

// Update implements IService.
func (s *service) Update(tourID string, dto *model.Tour) (*model.Tour, error) {
	panic("unimplemented")
}

func NewService(db *sql.DB, gorm *gorm.DB) IService {
	return &service{
		db:   db,
		gorm: gorm,
	}
}
