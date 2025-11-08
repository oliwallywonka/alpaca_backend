package rolefx

import "go.uber.org/fx"

var Module = fx.Module(
	"rolefx",
	fx.Provide(),
	fx.Invoke(),
)
