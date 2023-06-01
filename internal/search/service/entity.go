package service

import "database/sql"

type search struct {
	Data      string `json:"data,omitempty"`
	Login     string `json:"login" db:"login"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	LocalId   int    `json:"localId" db:"local_id"`
	IsFriends bool
}

type searchResult struct {
	Login     string         `db:"login"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	LocalId   int            `db:"local_id"`
}
