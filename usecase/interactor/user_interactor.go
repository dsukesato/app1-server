package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
)

type UsersInteractor struct {
	UsersRepository repository.UsersRepository
}

func (i *UsersInteractor) UserById(ctx context.Context, identifier int) (user model.User, err error) {
	user, err = i.UsersRepository.GetSelect(ctx, identifier)

	return
}

func (i *UsersInteractor) UsersAll(ctx context.Context) (users model.Users, err error) {
	users, err = i.UsersRepository.GetAll(ctx)

	return
}

// uRegistryはuser Registryの略
func (i *UsersInteractor) Add(ctx context.Context, uRequest model.PostUserRequest) (uRegistry model.User, err error) {
	id, err := i.UsersRepository.Store(ctx, uRequest)


	uRegistry, err = i.UsersRepository.GetSelect(ctx, id)

	return
}
