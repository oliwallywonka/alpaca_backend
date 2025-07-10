package userfx

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
	"github.com/oliwallywonka/alpaca_backend/internal/shared/utils"
)

type IService interface {
	GetPaginated(dto *shared.PaginatedQueryParamsDTO) (*shared.PaginatedResponse, error)
	GetByUniqueKey(key string) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(userID string, user *UpdateUserDTO) (*model.User, error)

	
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
	var users []struct {
		model.User
		Role model.Role
	}
	params := dto.DTOToModel()
	searchFilter := fmt.Sprintf("%%%s%%", params.SearchFilter)
	stm := User.SELECT(User.AllColumns, Role.Name, Role.Description).
		FROM(
			User.LEFT_JOIN(Role, Role.ID.EQ(User.RoleID)),
		).
		WHERE(OR(
			CAST(User.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(User.Email).AS_TEXT().LIKE(String(searchFilter)),
			CAST(User.Contacts).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Description).AS_TEXT().LIKE(String(searchFilter)),
		)).
		ORDER_BY(Raw(fmt.Sprintf("%s %s", params.OrderBy, params.OrderDirection))).
		LIMIT(params.Limit).
		OFFSET(params.Offset)
	err := stm.Query(s.db, &users)
	if err != nil {
		return nil, err
	}

	stmCount := SELECT(COUNT(STAR).AS("total")).
		FROM(User.LEFT_JOIN(Role, Role.ID.EQ(User.RoleID))).
		WHERE(OR(
			CAST(User.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(User.Email).AS_TEXT().LIKE(String(searchFilter)),
			CAST(User.Contacts).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Name).AS_TEXT().LIKE(String(searchFilter)),
			CAST(Role.Description).AS_TEXT().LIKE(String(searchFilter)),
		))
	err = stmCount.Query(s.db, &countStruct)
	if err != nil {
		return nil, err
	}
	var totalUsers int = 0
	if len(countStruct) > 0 {
		totalUsers = countStruct[0].Total
	}
	return &shared.PaginatedResponse{
		Items: users,
		Page:  int(dto.Page),
		Total: totalUsers,
	}, nil
}

// GetByUniqueKey implements IService.
func (s *service) GetByUniqueKey(key string) (*model.User, error) {
	stm := User.SELECT(User.AllColumns).
		FROM(User).
		WHERE(User.ID.EQ(String(key)))
	user := model.User{}
	err := stm.Query(s.db, &user)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements IService.
func (s *service) Save(user *model.User) (*model.User, error) {
	userExist, err := s.emailExists(user.Email)
	if userExist {
		return nil, errors.New("email already exists")
	}
	if err != nil {
		return nil, err
	}

	user.ID = uuid.NewString()
	if user.Password != "" {
		hash, err := utils.HashString(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hash
	}
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	stm := User.INSERT(User.AllColumns).
		MODEL(user).
		RETURNING(User.AllColumns)
	_, err = stm.Exec(s.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements IService.
func (s *service) Update(userID string, dto *UpdateUserDTO) (*model.User, error) {
	user := &model.User{}
	dto.UpdatedAt = int(time.Now().Unix())
	err := s.gorm.
		Model(user).
		Table("user").
		Where("id = ?", userID).
		Updates(dto).
		Clauses(clause.Returning{}).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) emailExists(email string) (bool, error) {
	var user model.User
	query := User.SELECT(User.ID).
		FROM(User).
		WHERE(User.Email.EQ(String(email)))
	err := query.Query(s.db, &user)
	if errors.Is(err, qrm.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if user.ID != "" {
		return true, nil
	}
	return false, nil
}
