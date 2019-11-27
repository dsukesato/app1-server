package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type PostsRepository interface {
	GetAll(context.Context) (model.Posts, error)
	GetSelect(context.Context, int) (model.Post, error)
	Store(context.Context, model.PostPostsRequest) (int, error)
}
