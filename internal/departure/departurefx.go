package departure

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Modulse = fx.Module(
	"departurefx",
	fx.Provide(

	),
	fx.Invoke(func(e *echo.Echo) {

	}),
)