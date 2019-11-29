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

type PostsController struct {
	Interactor usecase.PostsInteractor
}

func NewPostsController(dbconn database.DBConn) *PostsController {
	return &PostsController {
		Interactor: usecase.PostsInteractor {
			PostsRepository: &database.PostsRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *PostsController) PostsIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/posts/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	posts, err := c.Interactor.PostsAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *PostsController) PostsIdHandler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/Lookin/posts/{id}" {
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
	post, err := c.Interactor.PostsById(ctx, id)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type GetPostsResponse struct {
	Posts []PostsField `json:"posts"`
}

type PostsField struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	RestaurantId int       `json:"restaurant_id"`
	Image        string    `json:"image"`
	Good         int       `json:"good"`
	Genre        string    `json:"genre"`
	Comment      string    `json:"comment"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}

func (c *PostsController) PostsSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/posts/" {
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
	var jsonBody = new(model.PostPostsRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostPostsRequest{}
	request.UserId = jsonBody.UserId
	request.RestaurantId = jsonBody.RestaurantId
	request.Image = jsonBody.Image
	request.Genre = jsonBody.Genre
	request.Comment = jsonBody.Comment

	posts, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type PostPostsResponse struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	RestaurantId int       `json:"restaurant_id"`
	Image        string    `json:"image"`
	Good         int       `json:"good"`
	Genre        string    `json:"genre"`
	Comment      string    `json:"comment"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}
