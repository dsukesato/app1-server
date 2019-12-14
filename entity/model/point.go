package model

import "database/sql"

type Point struct {
	Id           int          `json:"id"`
	RestaurantId int          `json:"restaurant_id"`
	UserId       int          `json:"user_id"`
	Transaction  string       `json:"transaction"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

type Points []Point

type PostPointRequest struct {
	RestaurantId int          `json:"restaurant_id"`
	UserId       int          `json:"user_id"`
	Transaction  string       `json:"transaction"`
}

type PostPointResponse struct {
	Id           int          `json:"id"`
	RestaurantId int          `json:"restaurant_id"`
	UserId       int          `json:"user_id"`
	Transaction  string       `json:"transaction"`
	CreatedAt    sql.NullTime `json:"created_at"`
}
