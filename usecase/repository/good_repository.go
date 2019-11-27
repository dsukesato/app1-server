package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type GoodRepository interface {
	GetAll(context.Context) (model.Goods, error)
	GetSelect(context.Context, int) (model.Good, error)
	Store(context.Context, model.PostGoodRequest) (int, error)
}
