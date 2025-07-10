package jetfx

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: USE ENV VARIABLES
// TODO: GENERATE TABLES
// TODO: APPLY MIGRATIONS
func New() *sql.DB {
	db, err := sql.Open("sqlite3", "file:dev.db")
	if err != nil {
		log.Println("Failed to connect to DB:" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping DB:" + err.Error())
	}
	log.Println("Connected to DB")
	return db
}
