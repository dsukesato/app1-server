package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type RestaurantsRepository interface {
	GetAll(context.Context) (model.Restaurants, error)
	GetSelectById(context.Context, int) (model.Restaurant, error)
	GetLastId(context.Context) (int, error)
	Store(context.Context, model.RestaurantRequest) (int, error)
}
