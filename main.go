package main

import (
	"context"
	"log"

	"github.com/oliwallywonka/alpaca_backend/config"
	bunfx "github.com/oliwallywonka/alpaca_backend/pkg/bun"
	"go.uber.org/fx"
	/* pocketbasefx "github.com/oliwallywonka/alpaca_backend/pkg/pocketbase" */)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			config.New,
		),

		/* pocketbasefx.Module, */
		bunfx.Module,
	)
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
