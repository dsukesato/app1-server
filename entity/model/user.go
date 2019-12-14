package model

import (
	"database/sql"
)

type User struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Password  string       `json:"password"`
	Gender    string       `json:"gender"`
	BirthDay  string       `json:"birthday"`
	State     bool         `json:"state"`
	Point     int          `json:"point"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Users []User

type PostUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birthday"`
}

type PostUsersResponse struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Gender    string       `json:"gender"`
	BirthDay  string       `json:"birthday"`
	State     bool         `json:"state"`
	Point     int          `json:"point"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birthday"`
}

type SignUpResponse struct {
	Id       int    `json:"id"`
	Message  string `json:"message"`
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

type SignOutResponse struct {
	SignOutBool bool   `json:"sign_out_bool"`
	Message    string `json:"message"`
}
