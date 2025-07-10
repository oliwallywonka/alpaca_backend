package core

import (
	"fmt"

	"github.com/pocketbase/dbx"
	_ "modernc.org/sqlite"
)

func NewDB() (*dbx.DB, error) {
	db, err := dbx.Open("sqlite", "file:dev.db")

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	return db, nil
}
