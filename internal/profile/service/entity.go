package service

import (
	"database/sql"
	"github.com/gofrs/uuid"
)

type getFromDb struct {
	FirstNameDB   sql.NullString `db:"first_name"`
	LastNameDB    sql.NullString `db:"last_name"`
	AgeDB         sql.NullInt64  `db:"age"`
	GenderDB      sql.NullString `db:"gender"`
	DateOfBirthDB sql.NullString `db:"date_of_birth"`
}

type getUserData struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email,omitempty" db:"email"`
	Login       string    `json:"login,omitempty" db:"login"`
	FirstName   string    `json:"firstName,omitempty" db:"first_name"`
	LastName    string    `json:"lastName,omitempty" db:"last_name"`
	Age         int64     `json:"age,omitempty" db:"age"`
	Gender      string    `json:"gender,omitempty" db:"gender"`
	DateOfBirth string    `json:"date_of_birth" db:"date_of_birth"`
	LocalId     string
}

type editUserData struct {
	ID           string `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	Login        string `json:"login" db:"login"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName" db:"first_name"`
	LastName     string `json:"lastName" db:"last_name"`
	Age          int    `json:"age" db:"age"`
	Gender       string `json:"gender" db:"gender"`
	DateOfBirth  string `json:"date_of_birth" db:"date_of_birth"`
	HashPassword []byte `db:"password"`
	OldPassword  string `json:"oldPassword"`
	NewPassword  string `json:"newPassword"`
}
