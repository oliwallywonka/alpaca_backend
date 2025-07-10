package resourcefx

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type service struct {
	db   *sql.DB
	gorm *gorm.DB
}

type ResourceService interface {
	GetPaginated(dtoParams *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*model.Resource, error)
	Save(resource *model.Resource) (*model.Resource, error)
	Update(resourceID string, resource *UpdateResourceDTO) (*model.Resource, error)
}

func NewResourceService(db *sql.DB, gorm *gorm.DB) ResourceService {
	return &service{
		db:   db,
		gorm: gorm,
	}
}

// GetPaginated implements IResource.
func (s *service) GetPaginated(dtoParams *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	var countStruct []struct {
		Total int
	}
	var resources []struct {
		model.Resource
		resourceTypes []struct {
			model.ResourceResourceType
			resourceType model.ResourceType
		}
	}
	params := dtoParams.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := Resource.SELECT(Resource.AllColumns).
		FROM(Resource).
		WHERE(OR(
			CAST(Resource.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Resource.Description).AS_TEXT().LIKE(String(searchFilter)),
		)).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &resources)
	if err != nil {
		return nil, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Resource).
		WHERE(OR(
			CAST(Resource.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Resource.Description).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(s.db, &countStruct)
	if err != nil {
		return nil, err
	}
	var total int = 0
	if len(countStruct) > 0 {
		total = countStruct[0].Total
	}
	return &shared.PaginatedResponse{
		Items: resources,
		Page:  int(dtoParams.Page),
		Total: total,
	}, nil
}

// GetByUniqueKey implements IResource.
func (s *service) GetByUniqueKey(key string) (*model.Resource, error) {
	stm := Resource.SELECT(Resource.AllColumns).
		FROM(Resource).
		WHERE(Resource.ID.EQ(String(key))).
		LIMIT(1)
	resource := model.Resource{}
	err := stm.Query(s.db, &resource)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, fmt.Errorf("resource not found")
	}
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Save implements IResource.
func (s *service) Save(resource *model.Resource) (*model.Resource, error) {
	resource.ID = uuid.NewString()
	resource.CreatedAt = time.Now().Unix()
	resource.UpdatedAt = time.Now().Unix()
	stm := Resource.INSERT(Resource.AllColumns).
		MODEL(resource)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to save resource: %w", err)
	}
	return resource, nil
}

// Update implements IResource.
func (s *service) Update(resourceID string, dto *UpdateResourceDTO) (*model.Resource, error) {

	dto.UpdatedAt = int(time.Now().Unix())
	err := s.gorm.
		Table("resource").
		Where("id = ?", resourceID).
		Updates(dto).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to update resource: %w", err)
	}
	return nil, nil
}
