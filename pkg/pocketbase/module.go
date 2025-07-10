package pocketbasefx

import (
	"context"
	"log"

	"github.com/pocketbase/pocketbase"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"pocketbase",
	fx.Provide(
		pocketbase.New,
	),
	fx.Invoke(pbLifeHook),
)

func pbLifeHook(
	lifecycle fx.Lifecycle,
	p *pocketbase.PocketBase,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := p.Start(); err != nil {
				log.Fatal(err)
				return err
			}
			return nil
		},
	})
}
