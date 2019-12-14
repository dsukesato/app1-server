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
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PostUsersResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type SignInRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type SignInResponse struct {
	SignInBool bool   `json:"sign_in_bool"`
	Message    string `json:"message"`
}

type SignOutRequest struct {
	Id int `json:"id"`
}
