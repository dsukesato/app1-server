package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
)

type UsersRepository interface {
	GetAll(context.Context) (model.Users, error)
	GetSelect(context.Context, int) (model.User, error)
	Store(context.Context, model.PostUserRequest) (int, error)
}
