package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
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

func (i *PostsInteractor) PostsByRIG(ctx context.Context, rid int, genre string) (rig model.PRIG, err error) {
	rig, err = i.PostsRepository.GetSelectRIG(ctx, rid, genre)
	return
}

func (i *PostsInteractor) PostsAll(ctx context.Context) (posts model.Posts, err error) {
	posts, err = i.PostsRepository.GetAll(ctx)

	return
}

func (i *PostsInteractor) Add(ctx context.Context, posting model.PostsRequest) (posted model.Post, err error) {
	id, err := i.PostsRepository.Store(ctx, posting)

	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	posted, err = i.PostsRepository.GetSelect(ctx, id)

	return
}
