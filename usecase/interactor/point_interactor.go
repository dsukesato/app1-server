package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type PointInteractor struct {
	PointRepository repository.PointRepository
}

func (i *PointInteractor) PointById(ctx context.Context, identifier int) (point model.Point, err error) {
	point, err = i.PointRepository.GetSelect(ctx, identifier)

	return
}

func (i *PointInteractor) PointAll(ctx context.Context) (points model.Points, err error) {
	points, err = i.PointRepository.GetAll(ctx)

	return
}

func (i *PointInteractor) Add(ctx context.Context, poRequest model.PostPointRequest) (poRegistry model.Point, err error) {
	id, err := i.PointRepository.Store(ctx, poRequest)

	poRegistry, err = i.PointRepository.GetSelect(ctx, id)

	return
}