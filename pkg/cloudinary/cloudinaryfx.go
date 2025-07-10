package cloudinaryfx

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"cloudinaryfx",
	fx.Provide(
		NewService,
	),
)
