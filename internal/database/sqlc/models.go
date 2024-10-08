// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
	"time"
)

type Address struct {
	ID         string
	Cep        string
	Ibge       string
	Uf         string
	City       string
	Complement sql.NullString
	Street     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
