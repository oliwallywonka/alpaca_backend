package destinationfx

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type IService interface {
	GetAll(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByID(id string) (*model.Destination, error)
	Save(dto *model.Destination) (*model.Destination, error)
	Update(destinationID string, dto *model.Destination) (*model.Destination, error)
}

type service struct {
	db   *sql.DB
	gorm *gorm.DB
}

func NewService(db *sql.DB, gorm *gorm.DB) IService {
	return &service{
		db:   db,
		gorm: gorm,
	}
}

// GetAll returns all destinations
func (s *service) GetAll(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	var destinations []struct {
		model.Destination
		Parent model.Destination `alias:"parent"`
	}
	var countStruct []struct {
		Total int
	}
	params := dto.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	parent := Destination.AS("parent")
	stm := Destination.SELECT(Destination.AllColumns, parent.AllColumns.As("parent")).
		FROM(Destination.LEFT_JOIN(parent, parent.ID.EQ(Destination.ParentID))).
		WHERE(
			OR(
				CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Destination.Description).AS_TEXT().LIKE(String(searchFilter)),
				CAST(parent.Name).AS_TEXT().LIKE(String(searchFilter)),
			),
		).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &destinations)
	if err != nil {
		return nil, err
	}
	stmCount := Destination.SELECT(COUNT(STAR).AS("total")).
		FROM(Destination.LEFT_JOIN(parent, parent.ID.EQ(Destination.ParentID))).
		WHERE(
			OR(
				CAST(Destination.Name).AS_TEXT().LIKE(String(searchFilter)),
				CAST(Destination.Description).AS_TEXT().LIKE(String(searchFilter)),
				CAST(parent.Name).AS_TEXT().LIKE(String(searchFilter)),
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
		Items: destinations,
		Page:  int(dto.Page),
		Total: totalUsers,
	}, nil
}

// GetByID returns a destination by id
func (s *service) GetByID(id string) (*model.Destination, error) {
	var destination struct {
		model.Destination
		Parent model.Destination `alias:"parent"`
	}

	parent := Destination.AS("parent")
	stm := Destination.SELECT(Destination.AllColumns, parent.AllColumns).
		FROM(Destination.LEFT_JOIN(parent, parent.ID.EQ(Destination.ParentID))).
		WHERE(Destination.ID.EQ(String(id)))
	err := stm.Query(s.db, &destination)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.New("destination not found")
	}
	if err != nil {
		return nil, err
	}
	return &destination.Destination, nil
}

// Save creates a new destination
func (s *service) Save(dto *model.Destination) (*model.Destination, error) {
	dto.CreatedAt = time.Now().Unix()
	dto.UpdatedAt = time.Now().Unix()
	dto.ID = uuid.NewString()

	stm := Destination.INSERT(Destination.AllColumns).
		MODEL(dto).
		RETURNING(Destination.AllColumns)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, fmt.Errorf("error creating destination: %w", err)
	}
	return dto, nil
}

// Update updates a destination
func (s *service) Update(destinationID string, dto *model.Destination) (*model.Destination, error) {
	dto.UpdatedAt = time.Now().Unix()
	err := s.gorm.
		Table("destination").
		Where("id = ?", destinationID).
		Updates(dto).
		Error
	if err != nil {
		return nil, err
	}
	return dto, nil
}
