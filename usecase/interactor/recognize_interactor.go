package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type RecognizeInteractor struct {
	RecognizeRepository repository.RecognizeRepository
}

func (i *RecognizeInteractor) RecognizeById(ctx context.Context, identifier int) (recognize model.Recognize, err error) {
	recognize, err = i.RecognizeRepository.GetSelect(ctx, identifier)

	return
}

func (i *RecognizeInteractor) RecognizeAll(ctx context.Context) (rec model.Rec, err error) {
	rec, err = i.RecognizeRepository.GetAll(ctx)

	return
}

// uRegistryはuser Registryの略
func (i *RecognizeInteractor) Add(ctx context.Context, reRequest model.PostRecognizeRequest) (reRegistry model.Recognize, err error) {
	id, err := i.RecognizeRepository.Store(ctx, reRequest)


	reRegistry, err = i.RecognizeRepository.GetSelect(ctx, id)

	return
}
