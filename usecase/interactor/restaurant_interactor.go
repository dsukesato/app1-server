package interactor

import (
	"context"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
	"log"
)

type RestaurantsInteractor struct {
	RestaurantsRepository repository.RestaurantsRepository
}

func (i *RestaurantsInteractor) RestaurantLastId(ctx context.Context) (identifier int, err error) {
	identifier, err = i.RestaurantsRepository.GetLastId(ctx)
	return
}

func (i *RestaurantsInteractor) RestaurantById(ctx context.Context, identifier int) (restaurant model.Restaurant, err error) {
	restaurant, err = i.RestaurantsRepository.GetSelect(ctx, identifier)

	return
}

func (i *RestaurantsInteractor) RestaurantsAll(ctx context.Context) (restaurants model.Restaurants, err error) {
	restaurants, err = i.RestaurantsRepository.GetAll(ctx)

	return
}

// rRegistryはrestaurant Registryの略
func (i *RestaurantsInteractor) Add(ctx context.Context, rRequest model.RestaurantRequest) (rRegistry model.Restaurant, err error) {
	id, err := i.RestaurantsRepository.Store(ctx, rRequest)

	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	rRegistry, err = i.RestaurantsRepository.GetSelect(ctx, id)

	return
}

func (i *RestaurantsInteractor) Update(ctx context.Context, rRequest model.PutRestaurantRequest) (rResponse model.PutRestaurantResponse, err error) {
	restaurant, err := i.RestaurantsRepository.GetSelect(ctx, rRequest.Id)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	if restaurant.Name==rRequest.Name && restaurant.BusinessHours==rRequest.BusinessHours && restaurant.Image==rRequest.Image {
		log.Printf("以前のデータから更新された情報はありません\n")
	} else {
		id, err := i.RestaurantsRepository.Change(ctx, rRequest)
		if err != nil {
			log.Printf("err: %v\n", err)
		}
		log.Printf("データが更新されました\n")
		if rRequest.Id != id {
			rRes, err := i.RestaurantsRepository.GetSelect(ctx, id)
			if err != nil {
				log.Printf("err: %v\n", err)
			}
			rResponse = model.PutRestaurantResponse(rRes)
			err = fmt.Errorf("指定されたid: %dではなく、id: %dのデータを更新しました", rRequest.Id, id)
			return rResponse, err
		}
	}
	uRes, err := i.RestaurantsRepository.GetSelect(ctx, rRequest.Id)
	rResponse = model.PutRestaurantResponse(uRes)

	return
}
