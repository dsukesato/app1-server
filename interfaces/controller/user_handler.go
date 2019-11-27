package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type GetUsersResponse struct {
	Posts []PostsField `json:"posts"`
}

type UsersField struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Password     string       `json:"password"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

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

	user, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
