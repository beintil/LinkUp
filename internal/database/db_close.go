package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Close connect to database
func Close(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		panic(fmt.Errorf("error closed database connection: %s", err))
	}
}
