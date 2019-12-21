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

func (i *GoodInteractor) Add(ctx context.Context, request model.PostGoodRequest) (good model.PostGoodResponse, err error) {
	b := i.GoodRepository.GetSelectPUId(ctx, request.PostId, request.UserId)
	if b {
		id, err := i.GoodRepository.Store(ctx, request)
		if err != nil {
			log.Printf("err: %v with Store\n", err)
		}
		pId, _, err := i.GoodRepository.CountIncrease(ctx, request.PostId)
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
	} else {
		err = fmt.Errorf("もうすでにuser_id: %dはpost_id: %dにいいねしています", request.UserId, request.PostId)
		return
	}
	return
}

func (i *GoodInteractor) PutGood(ctx context.Context, request model.PutGoodRequest) (good model.PutGoodResponse, err error) {
	b := i.GoodRepository.GetSelectPUId(ctx, request.PostId, request.UserId)
	storeReq := model.PostGoodRequest{}
	storeReq.PostId = request.PostId
	storeReq.UserId = request.UserId
	var nGood int
	if b {
		id, err := i.GoodRepository.Store(ctx, storeReq)
		if err != nil {
			log.Printf("err: %v with Store\n", err)
		}
		_, nGood, err = i.GoodRepository.CountIncrease(ctx, request.PostId)
		if err != nil {
			log.Printf("err: %v with CountIncrease\n", err)
		}
		liked, err := i.GoodRepository.GetSelect(ctx, id)
		if err != nil {
			log.Printf("err: %v with GetSelect\n", err)
		}
		good.Id = id
		good.PostId = liked.PostId
		good.UserId = liked.UserId
		good.State = liked.State
		good.PostGood = nGood
	} else {
		id, err := i.GoodRepository.GetId(ctx, request.PostId, request.UserId)
		if err != nil {
			log.Printf("err: %v with GetId\n", err)
		}
		data, err := i.GoodRepository.GetSelect(ctx, id)
		if err != nil {
			log.Printf("err: %v with GetSelect\n", err)
		}
		if data.State == request.State {
			err = fmt.Errorf("以前のデータから変更がありません")
			return model.PutGoodResponse{}, err
		}
		state, err := i.GoodRepository.Change(ctx, request)
		if err != nil {
			log.Printf("err: %v with Change\n", err)
		}
		if !state {
			_, nGood, err = i.GoodRepository.CountDecrease(ctx, request.PostId)
			if err != nil {
				log.Printf("err: %v with CountDecrease\n", err)
			}
		} else {
			_, nGood, err = i.GoodRepository.CountIncrease(ctx, request.PostId)
			if err != nil {
				log.Printf("err: %v with CountDecrease\n", err)
			}
		}
		liked, err := i.GoodRepository.GetSelect(ctx, id)
		if err != nil {
			log.Printf("err: %v with GetSelect\n", err)
		}
		good.Id = id
		good.PostId = liked.PostId
		good.UserId = liked.UserId
		good.State = liked.State
		good.PostGood = nGood
	}
	return
}
