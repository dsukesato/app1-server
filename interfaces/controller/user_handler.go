package controller

import (
	"encoding/json"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"

	usecase "github.com/dsukesato/go13/pbl/app1-server/usecase/interactor"
)

type UsersController struct {
	Interactor usecase.UsersInteractor
}

func NewUsersController(dbconn database.DBConn) *UsersController {
	return &UsersController {
		Interactor: usecase.UsersInteractor {
			UsersRepository: &database.UsersRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *UsersController) UsersIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/users/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	posts, err := c.Interactor.UsersAll(ctx)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *UsersController) UsersIdHandler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/Lookin/users/{id}" {
	//	http.NotFound(w, r)
	//	return
	//}
	params := mux.Vars(r)
	log.Printf(params["id"])
	id, err := strconv.Atoi(params["id"])
	log.Printf("id: %d\n", id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	post, err := c.Interactor.UserById(ctx, id)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// SignUpHandlerと同じ機能
func (c *UsersController) UsersSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/users/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody = new(model.PostUserRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostUserRequest{}
	request.Name = jsonBody.Name
	request.Password = jsonBody.Password
	request.Gender = jsonBody.Gender
	request.BirthDay = jsonBody.BirthDay

	user, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *UsersController) UsersUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/users/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody = new(model.PutUserRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PutUserRequest{}
	request.Id = jsonBody.Id
	request.Name = jsonBody.Name
	request.Password = jsonBody.Password
	request.Gender = jsonBody.Gender
	request.BirthDay = jsonBody.BirthDay

	user, err := c.Interactor.Update(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *UsersController) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/sign_up/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody = new(model.SignUpRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.SignUpRequest{}
	request.Name = jsonBody.Name
	request.Password = jsonBody.Password
	request.Gender = jsonBody.Gender
	request.BirthDay = jsonBody.BirthDay

	signUp, err := c.Interactor.SignUp(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(signUp); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *UsersController) SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/sign_in/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody = new(model.SignInRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.SignInRequest{}
	request.Id = jsonBody.Id
	request.Password = jsonBody.Password

	// siBoolはsignInが成功しているかを判定するbool値
	response, err := c.Interactor.SignIn(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

//func (c *UsersController) SignOutHandler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/Lookin/sign_out/" {
//		http.NotFound(w, r)
//		return
//	}
//	ctx := r.Context()
//
//	if r.Header.Get("Content-Type") != "application/json" {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	//To allocate slice for request body
//	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	//Read body data to parse json
//	body := make([]byte, length)
//	length, err = r.Body.Read(body)
//	if err != nil && err != io.EOF {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	//parse json
//	var jsonBody = new(model.SignInRequest)
//	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	request := model.SignOutRequest{}
//	request.Id = jsonBody.Id
//
//	// siBoolはsignInが成功しているかを判定するbool値
//	siBool, err := c.Interactor.SignOut(ctx, request)
//
//	if err != nil {
//		log.Printf("err: %v\n", err)
//	}
//
//	response := model.SignInResponse{}
//	response.SignInBool = siBool
//	if siBool {
//		response.Message = "パスワード認証に成功しました"
//	} else {
//		response.Message = "パスワード認証に失敗しました"
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	if err = json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, "Internal Server Error", 500)
//		return
//	}
//}
