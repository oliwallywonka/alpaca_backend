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
)

type rpservice struct {
	db   *sql.DB
	gorm *gorm.DB
}

type ResourceProviderService interface {
	GetByProviderID(providerID string) (*[]ResourceProviderDTO, error)
	GetByID(providerResourceID string) (*ResourceProviderDTO, error)
	Save(resource *model.ResourceProvider) (*model.ResourceProvider, error)
	Update(providerResourceID string, resource *model.ResourceProvider) (*model.ResourceProvider, error)
	Delete(providerResourceID string) error
}

func NewResourceProviderService(db *sql.DB, gorm *gorm.DB) ResourceProviderService {
	return &rpservice{
		db:   db,
		gorm: gorm,
	}
}

// Save implements IResource.
func (s *rpservice) Save(ps *model.ResourceProvider) (*model.ResourceProvider, error) {
	ps.ID = uuid.NewString()
	ps.CreatedAt = time.Now().Unix()
	ps.UpdatedAt = time.Now().Unix()
	stm := ResourceProvider.INSERT(ResourceProvider.AllColumns).
		MODEL(ps)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to save provider resource: %w", err)
	}
	return ps, nil
}

// Update implements IResource.
func (s *rpservice) Update(psID string, ps *model.ResourceProvider) (*model.ResourceProvider, error) {
	ps.UpdatedAt = time.Now().Unix()
	err := s.gorm.
		Table("provider_resource").
		Where("id = ?", ps.ID).
		Updates(ps).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to update provider resource: %w", err)
	}
	return ps, nil
}

// Delete implements IResource.
func (s *rpservice) Delete(providerResourceID string) error {
	stm := ResourceProvider.DELETE().
		WHERE(ResourceProvider.ID.EQ(String(providerResourceID)))
	_, err := stm.Exec(s.db)
	if err != nil {
		return fmt.Errorf("failed to delete resource: %w", err)
	}
	return nil
}

// GetByProviderID implements IResource.
func (s *rpservice) GetByProviderID(providerID string) (*[]ResourceProviderDTO, error) {
	var providerResources []ResourceProviderDTO = make([]ResourceProviderDTO, 0)
	stm := ResourceProvider.SELECT(ResourceProvider.AllColumns, Provider.AllColumns, User.AllColumns, Resource.AllColumns).
		FROM(
			ResourceProvider.
				LEFT_JOIN(Provider, Provider.ID.EQ(ResourceProvider.ProviderID)).
				LEFT_JOIN(User, User.ID.EQ(ResourceProvider.UserID)).
				LEFT_JOIN(Resource, Resource.ID.EQ(ResourceProvider.ResourceID)),
		).
		WHERE(OR(
			ResourceProvider.ProviderID.EQ(String(providerID)),
			ResourceProvider.UserID.EQ(String(providerID)),
		))
	err := stm.Query(s.db, &providerResources)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return &providerResources, nil
		}
		return nil, fmt.Errorf("failed to get provider resources: %w", err)
	}
	return &providerResources, nil
}

// GetByID implements IResource.
func (s *rpservice) GetByID(providerResourceID string) (*ResourceProviderDTO, error) {
	var dto ResourceProviderDTO
	stm := ResourceProvider.SELECT(ResourceProvider.AllColumns, Provider.AllColumns, User.AllColumns, Resource.AllColumns).
		FROM(
			ResourceProvider.LEFT_JOIN(Provider, Provider.ID.EQ(ResourceProvider.ProviderID)),
			ResourceProvider.LEFT_JOIN(User, User.ID.EQ(ResourceProvider.UserID)),
			ResourceProvider.LEFT_JOIN(Resource, Resource.ID.EQ(ResourceProvider.ResourceID)),
		).
		WHERE(ResourceProvider.ID.EQ(String(providerResourceID)))
	err := stm.Query(s.db, &dto)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, fmt.Errorf("provider resource not found")
	}
	if err != nil {
		return nil, err
	}
	return &dto, nil
}
