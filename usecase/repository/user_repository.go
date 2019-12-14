package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type UsersRepository interface {
	GetAll(context.Context) (model.Users, error)
	GetSelect(context.Context, int) (model.User, error)
	GetPass(context.Context, int) (string, error)
	Store(context.Context, model.PostUserRequest) (int, error)
}
