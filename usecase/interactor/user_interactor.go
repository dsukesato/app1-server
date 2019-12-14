package interactor

import (
	"context"
	"fmt"
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

// uRegistryはuser Registryの略
func (i *UsersInteractor) Add(ctx context.Context, uRequest model.PostUserRequest) (uRegistry model.User, err error) {
	uRequest.Password, err = PasswordHash(uRequest.Password)
	id, err := i.UsersRepository.Store(ctx, uRequest)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	uRegistry, err = i.UsersRepository.GetSelect(ctx, id)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	return
}

func (i *UsersInteractor) SignUp(ctx context.Context, suReq model.SignUpRequest) (suRes model.SignUpResponse, err error) {
	suReq.Password, err = PasswordHash(suReq.Password)
	id, err := i.UsersRepository.Store(ctx, model.PostUserRequest(suReq))
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	uRegistry, err := i.UsersRepository.GetSelect(ctx, id)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	suRes.Id = uRegistry.Id
	suRes.Message = fmt.Sprintf("id: %dのユーザアカウントが作成されました", suRes.Id)

	return
}

// passwordのハッシュ化
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func (i *UsersInteractor) SignIn(ctx context.Context, request model.SignInRequest) (response model.SignInResponse, err error) {
	pass, err := i.UsersRepository.GetPass(ctx, request.Id)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	err = PasswordVerify(pass, request.Password)
	if err != nil {
		response.SignInBool = false
		response.Message = "パスワード認証に失敗しました"
	} else {
		response.SignInBool = true
		response.Message = "パスワード認証に成功しました"
	}

	return
}

// passwordの判定
func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

//func (i *UsersInteractor) SignOut(ctx context.Context, request model.SignOutRequest) (response model.SignOutResponse, err error) {
//	pass, err := i.UsersRepository.GetPass(ctx, request.Id)
//
//	err = PasswordVerify(pass, request.Password)
//	if err != nil {
//		return false, err
//	}
//
//	return true, nil
//}
