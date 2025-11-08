package rolefx

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles,alias:role"`

	ID          string                 `json:"id" bun:"id,pk,type:uuid"`
	Name        string                 `json:"name" bun:"name,notnull"`
	Description string                 `json:"description" bun:"description"`
	Permissions shared.PermissionField `json:"permissions" bun:"permissions,type:jsonb,default:[]"`
	IsActive    bool                   `json:"isActive" bun:"isActive,type:bool,default:true"`
	Created     time.Time              `json:"created" bun:"created,default:current_timestamp"`
	Updated     time.Time              `json:"updated" bun:"updated,default:current_timestamp"`
}
