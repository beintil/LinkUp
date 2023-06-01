package service

import "database/sql"

type users struct {
	Login     string `json:"login" db:"login"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	LocalId   int    `json:"localId" db:"local_id"`
}

type dataFromDb struct {
	Login     string         `db:"login"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	LocalId   int            `db:"local_id"`
}
