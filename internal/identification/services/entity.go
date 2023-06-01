package services

import "github.com/gofrs/uuid"

type User struct {
	ID           uuid.UUID `json:"id,omitempty" db:"id"`
	Email        string    `json:"email" db:"email"`
	Login        string    `json:"login" db:"login"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Password     string    `json:"password,omitempty"`
	HashPassword []byte    `db:"password" `
}
