package model

import (
	"database/sql"
)

type Post struct {
	Id           int          `json:"id"`
	UserId       int          `json:"user_id"`
	RestaurantId int          `json:"restaurant_id"`
	Content      string       `json:"content"`
	Good         int          `json:"good"`
	Genre        string       `json:"genre"`
	Comment      string       `json:"comment"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type Posts []Post

type PostsRequest struct {
	UserId       int    `json:"user_id"`
	RestaurantId int    `json:"restaurant_id"`
	Content      string `json:"content"`
	Genre        string `json:"genre"`
	Comment      string `json:"comment"`
}

type PostPostsRequest struct {
	UserId       int    `json:"user_id"`
	RestaurantId int    `json:"restaurant_id"`
	Genre        string `json:"genre"`
	Comment      string `json:"comment"`
}

type PostPostsResponse struct {
	Id           int          `json:"id"`
	UserId       int          `json:"user_id"`
	RestaurantId int          `json:"restaurant_id"`
	Content      string       `json:"content"`
	Good         int          `json:"good"`
	Genre        string       `json:"genre"`
	Comment      string       `json:"comment"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type PostsRIGResponse struct {
	Id           int          `json:"id"`
	Content      string       `json:"content"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

type PRIG []PostsRIGResponse
