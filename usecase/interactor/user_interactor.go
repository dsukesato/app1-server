package interactor

import (
	"context"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/usecase/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
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
func (i *UsersInteractor) Add(ctx context.Context, uRequest model.PostUserRequest) (uResponse model.PostUserResponse, err error) {
	//var check bool
	check, err := i.UsersRepository.CheckUuid(ctx, uRequest.Uuid)
	if err != nil{
		log.Printf("err: %v\n", err)
		//err = fmt.Errorf("%v", err)
	}
	if !check {
		err = fmt.Errorf("uuid is not unique. err: %v", err)
		return uResponse, err
	}
	uRequest.Password, err = PasswordHash(uRequest.Password)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	id, err := i.UsersRepository.Store(ctx, uRequest)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}

	uRegistry, err := i.UsersRepository.GetSelect(ctx, id)
	if err != nil {
		//log.Printf("err: %v\n", err)
		return
	}
	uResponse = model.PostUserResponse(uRegistry)

	return
}

func (i *UsersInteractor) Update(ctx context.Context, request model.PutUserRequest) (uResponse model.PutUserResponse, err error) {
	user, err := i.UsersRepository.GetSelect(ctx, request.Id)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	if err = PasswordVerify(user.Password, request.Password); err != nil {
		request.Password, err = PasswordHash(request.Password)
		if err != nil {
			log.Printf("err: %v\n", err)
		}
	} else {
		log.Printf("passwordは変更しません\n")
		request.Password = user.Password
	}
	tz := "T00:00:00Z"
	birthday := fmt.Sprintf("%s%s", request.BirthDay, tz)
	if user.Uuid==request.Uuid && user.Name==request.Name && user.Password==request.Password && user.Gender==request.Gender && user.BirthDay==birthday {
		log.Printf("以前のデータから更新された情報はありません\n")
	} else {
		id, err := i.UsersRepository.Change(ctx, request)
		if err != nil {
			log.Printf("err: %v\n", err)
		}
		log.Printf("データが更新されました\n")
		if request.Id != id {
			uRes, err := i.UsersRepository.GetSelect(ctx, id)
			if err != nil {
				log.Printf("err: %v\n", err)
			}
			uResponse = model.PutUserResponse(uRes)
			err = fmt.Errorf("指定されたid: %dではなく、id: %dのデータを更新しました", request.Id, id)
			return uResponse, err
		}
	}
	uRes, err := i.UsersRepository.GetSelect(ctx, request.Id)
	uResponse = model.PutUserResponse(uRes)

	return
}

func (i *UsersInteractor) SignUp(ctx context.Context, suReq model.SignUpRequest) (suRes model.SignUpResponse, err error) {
	check, err := i.UsersRepository.CheckUuid(ctx, suReq.Uuid)
	if err != nil{
		err = fmt.Errorf("%v", err)
	}
	if !check {
		err = fmt.Errorf("uuid is not unique. err: %v", err)
		return
	}
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
	identifier, err := i.UsersRepository.GetSelectUuid(ctx, request.Uuid)
	pass, err := i.UsersRepository.GetPass(ctx, identifier)
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
