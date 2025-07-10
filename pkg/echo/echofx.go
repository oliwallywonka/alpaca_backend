package echofx

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	"github.com/oliwallywonka/alpaca_backend/settings"
)

var Module = fx.Module(
	"echofx",
	fx.Provide(
		echo.New,
	),
	fx.Invoke(echoLifeHook),
)

func echoLifeHook(lifecycle fx.Lifecycle, e *echo.Echo, s *settings.Settings) {

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			//e.Use(middleware.Logger())
			e.Use(middleware.Recover())

			e.HTTPErrorHandler = shared.HttpErrorHandler

			corsConfig := middleware.CORSConfig{
				AllowOrigins: strings.Split(s.AllowedOrigins, ","),
				AllowMethods: strings.Split(s.AllowedMethods, ","),
			}
			e.Use(middleware.CORSWithConfig(corsConfig))

			go func() {
				if err := e.Start(":" + s.Port); err != nil {
					e.Logger.Fatal("shutting down the server: ", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
