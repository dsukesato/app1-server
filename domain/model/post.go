package model

import (
	"database/sql"
)

type Post struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	RestaurantId int       `json:"restaurant_id"`
	Image        string    `json:"image"`
	Good         int       `json:"good"`
	Genre        string    `json:"genre"`
	Comment      string    `json:"comment"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type Posts []Post

type PostPostsRequest struct {
	UserId       int    `json:"user_id"`
	RestaurantId int    `json:"restaurant_id"`
	Image        string `json:"image"`
	Genre        string `json:"genre"`
	Comment      string `json:"comment"`
}
