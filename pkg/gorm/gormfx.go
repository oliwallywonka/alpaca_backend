package gormfx

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/oliwallywonka/alpaca_backend/settings"
)

func New(s *settings.Settings) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file:dev.db"))

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
