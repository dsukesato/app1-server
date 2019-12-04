package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type PostsRepository interface {
	GetLastId(context.Context) (int, error)
	GetSelect(context.Context, int) (model.Post, error)
	GetSelectRIG(context.Context, int, string) (model.PRIG, error)
	GetAll(context.Context) (model.Posts, error)
	Store(context.Context, model.PostsRequest) (int, error)
}
