package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
	"log"
)

type RecognizeInteractor struct {
	RecognizeRepository repository.RecognizeRepository
}

func (i *RecognizeInteractor) RecognizeById(ctx context.Context, identifier int) (recognize model.Recognize, err error) {
	recognize, err = i.RecognizeRepository.GetSelect(ctx, identifier)

	return
}

func (i *RecognizeInteractor) RecognizeByUId(ctx context.Context, uid int) (rrs model.RecResponse, err error) {
	rids, err := i.RecognizeRepository.GetSelectUID(ctx, uid)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	for _, rid := range rids {
		rr, err := i.RecognizeRepository.GetSelectRID(ctx, rid)
		if err != nil {
			log.Printf("err: %v\n", err)
		}
		rrs = append(rrs, rr)
	}

	return
}

func (i *RecognizeInteractor) RecognizeAll(ctx context.Context) (rec model.Rec, err error) {
	rec, err = i.RecognizeRepository.GetAll(ctx)

	return
}

func (i *RecognizeInteractor) Add(ctx context.Context, reRequest model.PostRecognizeRequest) (reRegistry model.Recognize, err error) {
	id, err := i.RecognizeRepository.Store(ctx, reRequest)

	reRegistry, err = i.RecognizeRepository.GetSelect(ctx, id)

	return
}
