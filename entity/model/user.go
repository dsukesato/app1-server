package model

import (
	"database/sql"
)

type User struct {
	Id        int          `json:"id"`
	Uuid      string       `json:"uuid"`
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
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birthday"`
}

type PostUserResponse struct {
	Id        int          `json:"id"`
	Uuid      string       `json:"uuid"`
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

type PutUserRequest struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birthday"`
}

type PutUserResponse struct {
	Id        int          `json:"id"`
	Uuid      string       `json:"uuid"`
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

type SignUpRequest struct {
	Uuid     string `json:"uuid"`
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
	Uuid     string `json:"uuid"`
	Password string `json:"password"`
}

type SignInResponse struct {
	SignInBool bool   `json:"sign_in_bool"`
	Message    string `json:"message"`
}

//type SignOutRequest struct {
//	Uuid string `json:"uuid"`
//}
//
//type SignOutResponse struct {
//	SignOutBool bool  `json:"sign_out_bool"`
//	Message    string `json:"message"`
//}
