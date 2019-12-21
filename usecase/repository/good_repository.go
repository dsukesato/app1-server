package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type GoodRepository interface {
	GetAll(context.Context) (model.Goods, error)
	GetId(context.Context, int, int) (int, error)
	GetGood(context.Context, int) (int, error)
	GetSelect(context.Context, int) (model.Good, error)
	GetSelectPUId(context.Context, int ,int) bool
	Store(context.Context, model.PostGoodRequest) (int, error)
	Change(context.Context, model.PutGoodRequest) (bool, error)
	CountIncrease(context.Context, int) (int, int, error)
	CountDecrease(context.Context, int) (int, int, error)
}
