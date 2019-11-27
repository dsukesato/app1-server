package model

import (
	"database/sql"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Users []User

type PostUserRequest struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type PostUsersResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}
