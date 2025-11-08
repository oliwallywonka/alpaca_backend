package resourcefx

import (
	"time"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/uptrace/bun"
)

type Resource struct {
	bun.BaseModel `bun:"table:resources,alias:resource"`

	ID            string               `json:"id" bun:"id,pk,type:uuid"`
	DestinationID string               `json:"destinationId" bun:"destinationId"`
	Name          shared.LangField     `json:"name" bun:"name,type:jsonb,default:'{}'"`
	Images        shared.ImageField    `json:"images" bun:"images,type:jsonb,default:'[]'"`
	IsActive      bool                 `json:"isActive" bun:"isActive,notnull,default:true"`
	Types         []string             `json:"types" bun:"types,type:jsonb,default:'[]'"`
	Location      shared.LocationField `json:"location" bun:"location,type:jsonb,default:'{}'"`
	Created       time.Time            `json:"created" bun:"created,default:current_timestamp"`
	Updated       time.Time            `json:"updated" bun:"updated,default:current_timestamp"`
}

type ResourceProviders struct {

}
