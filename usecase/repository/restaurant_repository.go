package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type RestaurantsRepository interface {
	GetAll(context.Context) (model.Restaurants, error)
	GetSelectById(context.Context, int) (model.Restaurant, error)
	Store(context.Context, model.PostRestaurantRequest) (int, error)
}
