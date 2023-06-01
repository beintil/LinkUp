package database

import (
	"LinkUp_Update/config"
	"LinkUp_Update/var/logs"
	"context"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

// Init takes a "context.Context" parameter and returns a pointer to a "DB" object.
// It initializes a new "DB" object with a background context and calls the "dbDefinition" function.
func Init(ctx context.Context) *DB {
	db := &DB{
		ctx: ctx,
	}
	db.dbDefinition()
	return db
}

// GetConn returns the "conn"
func (db *DB) GetConn() *sqlx.DB {
	return db.conn
}

// The dbDefinition function takes a "config.Config" and a pointer to a "DB" object.
// It checks the value of the "ConnectDB" field in the configuration and calls the corresponding connection method.
func (db *DB) dbDefinition() {
	switch strings.ToLower(config.Get("CONNECT_DB").ToString()) {
	case "postgres":
		db.connectionPSQL()
	case "mongodb":
	// ...
	default:
		logs.Get().LogApi(errors.New("There is currently no connection for the "+config.Get("CONNECT_DB").ToString()+" database"), func() {
			log.Fatal()
		})
	}
}

// The "connection" method creates a PostgreSQL connection string using the configuration settings and attempts
// to connect to the database using the "pgx" driver and "sqlx" library. If the connection is successful,
// the "conn" field of the "DB" object is set to the connected database.
func (db *DB) connectionPSQL() {
	defer func() {
		if rec := recover(); rec != nil {
			db.log.LogApi(errors.New(fmt.Sprint(rec)), func() {
				log.Fatal()
			})
		}
	}()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Get("DATABASE_USERNAME").ToString(),
		config.Get("DATABASE_PASSWORD").ToString(),
		config.Get("DATABASE_HOST").ToString(),
		config.Get("DATABASE_PORT").ToString(),
		config.Get("DATABASE_NAME").ToString(),
		config.Get("SSLMODE").ToString(),
	)
	conn, err := sqlx.ConnectContext(db.ctx, "pgx", dsn)
	if err != nil {
		db.log.LogApi(fmt.Errorf("failed to connect to database: %v", err), func() {
			log.Fatal()
		})
	}
	db.conn = conn
}
