package userfx

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/internal/rolefx"
	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:user"`

	ID       string                `json:"id" bun:"id,pk,type:uuid"`
	RoleID   string                `json:"roleId" bun:"roleId,notnull"`
	Email    string                `json:"email" bun:"email,notnull,unique,type:varchar(255)"`
	Password string                `json:"password" bun:"password,notnull,type:varchar(255)"`
	Name     string                `json:"name" bun:"name,notnull,type:varchar(255)"`
	Avatar   string                `json:"avatar" bun:"avatar,notnull,type:varchar(255)"`
	Contacts []shared.ContactField `json:"contacts" bun:"contacts,type:jsonb,default:[]"`
	IsActive bool                  `json:"isActive" bun:"isActive,notnull,default:true"`
	Created  time.Time             `json:"created" bun:"created,default:current_timestamp"`
	Updated  time.Time             `json:"updated" bun:"updated,default:current_timestamp"`
	Role     rolefx.Role           `json:"role,omitempty" bun:"rel:belongs-to,join:roleId=id"`
}
