package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type RestaurantsRepository interface {
	GetAll(context.Context) (model.Restaurants, error)
	GetSelect(context.Context, int) (model.Restaurant, error)
	//GetSelectUuid(context.Context, string) (int, error)
	CheckUuid(context.Context, string) (bool, error)
	GetLastId(context.Context) (int, error)
	Store(context.Context, model.RestaurantRequest) (int, error)
	Change(context.Context, model.PutRestaurantRequest) (int, error)
}
