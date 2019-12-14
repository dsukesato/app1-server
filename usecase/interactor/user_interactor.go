package interactor

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
	"golang.org/x/crypto/bcrypt"
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

// SignUpと同じ機能
// uRegistryはuser Registryの略
func (i *UsersInteractor) Add(ctx context.Context, uRequest model.PostUserRequest) (uRegistry model.User, err error) {
	id, err := i.UsersRepository.Store(ctx, uRequest)

	uRegistry, err = i.UsersRepository.GetSelect(ctx, id)

	return
}

func (i *UsersInteractor) SignUp(ctx context.Context, uRequest model.PostUserRequest) (uRegistry model.User, err error) {
	id, err := i.UsersRepository.Store(ctx, uRequest)

	uRegistry, err = i.UsersRepository.GetSelect(ctx, id)

	return
}

func (i *UsersInteractor) SignIn(ctx context.Context, request model.SignInRequest) (bool, error) {
	pass, err := i.UsersRepository.GetPass(ctx, request.Id)

	err = PasswordVerify(pass, request.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

// passwordの判定
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
