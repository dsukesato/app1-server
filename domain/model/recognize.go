package model

import (
	"database/sql"
)

type Recognize struct {
	Id           int       `json:"id"`
	RestaurantId int       `json:"restaurant_id"`
	UserId       int       `json:"user_id"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

type Rec []Recognize

type PostRecognizeRequest struct {
	RestaurantId int `json:"restaurant_id"`
	UserId       int `json:"user_id"`
}

type RecognizeResponse struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Image     string       `json:"image"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type RecResponse []RecognizeResponse
