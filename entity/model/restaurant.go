package model

import (
	"database/sql"
)

type Restaurant struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	BusinessHours string       `json:"business_hours"`
	Image         string       `json:"image"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
}

type Restaurants []Restaurant

type RestaurantRequest struct {
	Name          string `json:"name"`
	BusinessHours string `json:"business_hours"`
	Image         string `json:"image"`
}

type PostRestaurantRequest struct {
	Name          string `json:"name"`
	BusinessHours string `json:"business_hours"`
}

type PostRestaurantResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	BusinessHours string `json:"business_hours"`
	Image         string `json:"image"`
}

type PutRestaurantRequest struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	BusinessHours string `json:"business_hours"`
	Image         string `json:"image"`
}

type PutRestaurantResponse struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	BusinessHours string       `json:"business_hours"`
	Image         string       `json:"image"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
}
