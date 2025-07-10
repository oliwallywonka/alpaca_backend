package main

import (
	//jetfx "github.com/oliwallywonka/alpaca_backend/pkg/jet"
	"context"
	"log"

	"go.uber.org/fx"

	pocketbasefx "github.com/oliwallywonka/alpaca_backend/pkg/pocketbase"
	/* destinationfx "github.com/oliwallywonka/alpaca_backend/internal/destination"
	providerfx "github.com/oliwallywonka/alpaca_backend/internal/provider"
	resourcefx "github.com/oliwallywonka/alpaca_backend/internal/resource"
	rolefx "github.com/oliwallywonka/alpaca_backend/internal/role"
	tourfx "github.com/oliwallywonka/alpaca_backend/internal/tour"
	userfx "github.com/oliwallywonka/alpaca_backend/internal/user"
	"github.com/oliwallywonka/alpaca_backend/pkg/cloudinary"
	"github.com/oliwallywonka/alpaca_backend/pkg/echo"
	"github.com/oliwallywonka/alpaca_backend/pkg/gorm"
	"github.com/oliwallywonka/alpaca_backend/pkg/jet"
	"github.com/oliwallywonka/alpaca_backend/settings" */)

func main() {
	//jetfx.Generate()
	app := fx.New(
		fx.Provide(
			context.Background,
			/* settings.New,
			gormfx.New,
			jetfx.New, */
		),
		// INTERNAL MODULES
		/* rolefx.Module,
		userfx.Module,
		providerfx.Module,
		resourcefx.Module,
		destinationfx.Module,
		tourfx.Module,
		// CUSTOM MODULES
		echofx.Module,
		cloudinaryfx.Module, */

		pocketbasefx.Module,
	)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
