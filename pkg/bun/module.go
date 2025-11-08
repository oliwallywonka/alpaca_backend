package bunfx

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/fx"

	"github.com/oliwallywonka/alpaca_backend/config"
	"github.com/oliwallywonka/alpaca_backend/internal/resourcefx"
	"github.com/oliwallywonka/alpaca_backend/internal/rolefx"
	"github.com/oliwallywonka/alpaca_backend/internal/userfx"
)

func New(ctx context.Context, conf *config.Config, lc fx.Lifecycle) *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, conf.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
	return db
}

func insertTables(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*rolefx.Role)(nil),
		(*userfx.User)(nil),
		(*resourcefx.Resource)(nil),
	}
	for _, model := range models {
		_, err := db.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

var Module = fx.Module(
	"bunfx",
	fx.Provide(New),
	fx.Invoke(insertTables),
)
