package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type PointRepository interface {
	GetAll(context.Context) (model.Points, error)
	GetSelect(context.Context, int) (model.Point, error)
	Store(context.Context, model.PostPointRequest) (int, error)
}
