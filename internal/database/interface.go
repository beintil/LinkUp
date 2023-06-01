package database

import (
	"LinkUp_Update/var/logs"
	"context"
	"github.com/jmoiron/sqlx"
)

type database interface {
	GetConn() *sqlx.DB
	connectionPSQL()
	dbDefinition()
}

// DB with three fields: "conn" of type
// "*sqlx.DB" (a connection to a database), "tx" of type "*sqlx.Tx" (a database transaction),
// and "ctx" of type "context.Context" (a context for the database connection).
type DB struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
	ctx  context.Context
	log  *logs.Logging
}
