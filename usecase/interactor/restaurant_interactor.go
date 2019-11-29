package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type RestaurantsInteractor struct {
	RestaurantsRepository repository.RestaurantsRepository
}

func (i *RestaurantsInteractor) RestaurantById(ctx context.Context, identifier int) (restaurant model.Restaurant, err error) {
	restaurant, err = i.RestaurantsRepository.GetSelectById(ctx, identifier)

	return
}

func (i *RestaurantsInteractor) RestaurantsAll(ctx context.Context) (restaurants model.Restaurants, err error) {
	restaurants, err = i.RestaurantsRepository.GetAll(ctx)

	return
}

// rRegistryはrestaurant Registryの略
func (i *RestaurantsInteractor) Add(ctx context.Context, rRequest model.PostRestaurantRequest) (rRegistry model.Restaurant, err error) {
	id, err := i.RestaurantsRepository.Store(ctx, rRequest)

	rRegistry, err = i.RestaurantsRepository.GetSelectById(ctx, id)

	return
}
