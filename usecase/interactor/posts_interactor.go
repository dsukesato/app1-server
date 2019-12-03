package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type PostsInteractor struct {
	PostsRepository repository.PostsRepository
}

func (i *PostsInteractor) PostsLastId(ctx context.Context) (identifier int, err error) {
	identifier, err = i.PostsRepository.GetLastId(ctx)
	return
}

func (i *PostsInteractor) PostsById(ctx context.Context, identifier int) (post model.Post, err error) {
	post, err = i.PostsRepository.GetSelect(ctx, identifier)

	return
}

func (i *PostsInteractor) PostsAll(ctx context.Context) (posts model.Posts, err error) {
	posts, err = i.PostsRepository.GetAll(ctx)

	return
}

func (i *PostsInteractor) Add(ctx context.Context, posting model.PostsRequest) (posted model.Post, err error) {
	id, err := i.PostsRepository.Store(ctx, posting)

	posted, err = i.PostsRepository.GetSelect(ctx, id)

	return
}
