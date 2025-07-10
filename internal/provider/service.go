package providerfx

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/google/uuid"
	"github.com/oliwallywonka/alpaca_backend/db/model"
	. "github.com/oliwallywonka/alpaca_backend/db/table"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
)

type IService interface {
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*model.Provider, error)
	Save(provider *model.Provider) (*model.Provider, error)
	Update(providerID string, provider *UpdateProviderDTO) (*model.Provider, error)
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

func (s *service) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	var countStruct []struct {
		Total int
	}
	var providers []model.Provider
	params := dto.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := Provider.SELECT(Provider.AllColumns).
		FROM(Provider).
		WHERE(OR(
			CAST(Provider.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Provider.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Provider.Contacts).AS_TEXT().LIKE(String(searchFilter)),
		)).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &providers)
	if err != nil {
		return nil, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Provider).
		WHERE(OR(
			CAST(Provider.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Provider.Description).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Provider.Contacts).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(s.db, &countStruct)
	if err != nil {
		return nil, err
	}
	var totalProviders int = 0
	if len(countStruct) > 0 {
		totalProviders = countStruct[0].Total
	}
	return &shared.PaginatedResponse{
		Items: providers,
		Page:  int(dto.Page),
		Total: totalProviders,
	}, nil
}

// GetByUniqueKey implements IService.
func (s *service) GetByUniqueKey(key string) (*model.Provider, error) {
	stm := Provider.SELECT(Provider.AllColumns).
		FROM(Provider).
		WHERE(Provider.ID.EQ(String(key)))
	provider := model.Provider{}
	err := stm.Query(s.db, &provider)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.New("provider not found")
	}
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

// Save implements IService.
func (s *service) Save(provider *model.Provider) (*model.Provider, error) {
	provider.ID = uuid.NewString()
	provider.CreatedAt = time.Now().Unix()
	provider.UpdatedAt = time.Now().Unix()

	stm := Provider.INSERT(Provider.AllColumns).
		MODEL(provider).
		RETURNING(Provider.AllColumns)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

// Update implements IService.
func (s *service) Update(providerID string, dto *UpdateProviderDTO) (*model.Provider, error) {
	provider := &model.Provider{}
	dto.UpdatedAt = int(time.Now().Unix())
	err := s.gorm.
		Model(provider).
		Table("provider").
		Where("id = ?", providerID).
		Updates(dto).
		Clauses(clause.Returning{}).
		Error
	if err != nil {
		return nil, err
	}
	return provider, nil
}
