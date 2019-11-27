package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type RecognizeRepository interface {
	GetAll(context.Context) (model.Rec, error)
	GetSelect(context.Context, int) (model.Recognize, error)
	Store(context.Context, model.PostRecognizeRequest) (int, error)
}
