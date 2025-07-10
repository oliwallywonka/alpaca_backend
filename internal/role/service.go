package rolefx

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
	GetByUniqueKey(key string) (*model.Role, error)
	Save(role *model.Role) (*model.Role, error)
	Update(roleID string, role *UpdateRoleDTO) (*model.Role, error)
}

type roleService struct {
	db   *sql.DB
	gorm *gorm.DB
}

func NewService(db *sql.DB, gorm *gorm.DB) IService {
	return &roleService{
		db:   db,
		gorm: gorm,
	}
}

func (s *roleService) GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error) {
	var countStruct []struct {
		Total int
	}
	var roles []model.Role
	params := dto.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := Role.SELECT(Role.AllColumns).
		FROM(Role).
		WHERE(OR(
			CAST(Role.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Description).AS_TEXT().LIKE(String(searchFilter)),
		)).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &roles)
	if err != nil {
		return nil, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(Role).
		WHERE(OR(
			CAST(Role.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Description).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(s.db, &countStruct)
	if err != nil {
		return nil, err
	}
	var totalRoles int = 0
	if len(countStruct) > 0 {
		totalRoles = countStruct[0].Total
	}
	return &shared.PaginatedResponse{
		Items: roles,
		Page:  int(dto.Page),
		Total: totalRoles,
	}, nil
}

// GetByUniqueKey implements IService.
func (s *roleService) GetByUniqueKey(key string) (*model.Role, error) {
	stm := Role.SELECT(Role.AllColumns).
		FROM(Role).
		WHERE(Role.ID.EQ(String(key)))
	role := model.Role{}
	err := stm.Query(s.db, &role)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.New("role not found")
	}
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// Save implements IService.
func (s *roleService) Save(role *model.Role) (*model.Role, error) {
	role.ID = uuid.New().String()
	role.CreatedAt = time.Now().Unix()
	role.UpdatedAt = time.Now().Unix()
	stm := Role.INSERT(Role.AllColumns).
		MODEL(role).
		RETURNING(Role.AllColumns)
	_, err := stm.Exec(s.db)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// Update implements IService.
func (s *roleService) Update(roleID string, dto *UpdateRoleDTO) (*model.Role, error) {
	role := &model.Role{}
	dto.UpdatedAt = int(time.Now().Unix())
	err := s.gorm.
		Model(role).
		Table("role").
		Where("id = ?", roleID).
		Updates(dto).
		Clauses(clause.Returning{}).
		Error
	if err != nil {
		return nil, err
	}
	return role, nil
}
