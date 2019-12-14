package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type GoodInteractor struct {
	GoodRepository repository.GoodRepository
}

func (i *GoodInteractor) GoodById(ctx context.Context, identifier int) (good model.Good, err error) {
	good, err = i.GoodRepository.GetSelect(ctx, identifier)

	return
}

func (i *GoodInteractor) GoodAll(ctx context.Context) (goods model.Goods, err error) {
	goods, err = i.GoodRepository.GetAll(ctx)

	return
}

func (i *GoodInteractor) Add(ctx context.Context, liking model.PostGoodRequest) (liked model.Good, err error) {
	id, err := i.GoodRepository.Store(ctx, liking)

	liked, err = i.GoodRepository.GetSelect(ctx, id)

	return
}
