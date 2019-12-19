package interactor

import (
	"context"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
	"log"
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

func (i *GoodInteractor) Add(ctx context.Context, liking model.PostGoodRequest) (good model.PostGoodResponse, err error) {
	b := i.GoodRepository.GetSelectPUId(ctx, liking.PostId, liking.UserId)

	if b {
		id, err := i.GoodRepository.Store(ctx, liking)
		if err != nil {
			log.Printf("err: %v with Store\n", err)
		}
		pId, nGood, err := i.GoodRepository.CountIncrease(ctx, liking.PostId)
		if err != nil {
			log.Printf("err: %v with CountIncrease\n", err)
		}
		liked, err := i.GoodRepository.GetSelect(ctx, id)
		if err != nil {
			log.Printf("err: %v with GetSelect\n", err)
		}
		good.Id = id
		good.PostId = pId
		good.UserId = liked.UserId
		good.State = liked.State
		good.PostGood = nGood
	} else {
		err = fmt.Errorf("もうすでにuser_id: %dはpost_id: %dにいいねしています", liking.UserId, liking.PostId)
		return
	}
	return
}
