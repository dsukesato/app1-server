package repository

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
)

type UsersRepository interface {
	GetAll(context.Context) (model.Users, error)
	GetSelect(context.Context, int) (model.User, error)
	GetSelectUuid(context.Context, string) (int, error)
	CheckUuid(context.Context, string) (bool, error)
	GetPass(context.Context, int) (string, error)
	Store(context.Context, model.PostUserRequest) (int, error)
	Change(context.Context, model.PutUserRequest) (int, error)
}
