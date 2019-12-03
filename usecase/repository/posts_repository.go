package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type PostsRepository interface {
	GetAll(context.Context) (model.Posts, error)
	GetSelect(context.Context, int) (model.Post, error)
	GetLastId(context.Context) (int, error)
	Store(context.Context, model.PostsRequest) (int, error)
}
