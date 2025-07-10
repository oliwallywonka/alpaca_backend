package rolefx

import "github.com/oliwallywonka/alpaca_backend/internal/shared"

type UpdateRoleDTO struct {
	Name        string
	Description string
	// GORM DOESNT UPDATE ZERO VALUE FIELDS SO ITS MANDATORY TO USE POINTERS TO IGNORE THAT
	IsActive    *bool
	Permissions shared.PermissionField
	UpdatedAt   int
}
