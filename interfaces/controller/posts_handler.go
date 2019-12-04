package controller

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
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

func (c *PostsController) PostsRIGHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["restaurant_id"])
	rid, err := strconv.Atoi(params["restaurant_id"])
	log.Printf("restaurant_id: %d\n", rid)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	genre := params["genre"]
	log.Printf("genre: %s\n", genre)

	ctx := r.Context()
	post, err := c.Interactor.PostsByRIG(ctx, rid, genre)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *PostsController) PostsSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/posts/" {
		http.NotFound(w, r)
		return
	}

	//if r.Header.Get("Content-Type") != "application/json" {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	formValue := r.FormValue("json")

	var jsonBody model.PostPostsRequest

	b := []byte(formValue)
	err := json.Unmarshal(b, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	if jsonBody.Genre != "mood" && jsonBody.Genre != "food" && jsonBody.Genre != "drink" && jsonBody.Genre != "dessert" {
		log.Printf("Not the desired request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	formFile, _, err := r.FormFile("image")
	handleError(err)
	defer formFile.Close()

	ctx := r.Context()
	lastId, err := c.Interactor.PostsLastId(ctx)

	bucket := "pbl-lookin-storage" // GCSバケット名
	obj := fmt.Sprintf("post%d_%s.jpeg", lastId+1, jsonBody.Genre)
	bCtx := context.Background()

	client, err := storage.NewClient(bCtx)
	if err != nil {
		log.Printf("failed to create gcs client : %v", err)
	}

	// GCS writer
	writer := client.Bucket(bucket).Object(obj).NewWriter(bCtx)
	writer.ContentType = "image/jpeg" // 任意のContentTypeに置き換える

	// uploadされた画像をgcsのwriterにコピー
	_, err = io.Copy(writer, formFile)
	handleError(err)

	if err := writer.Close(); err != nil {
		log.Printf("failed to close gcs writer : %v", err)
	}
	w.WriteHeader(http.StatusCreated)

	request := model.PostsRequest{}
	request.RestaurantId = jsonBody.RestaurantId
	request.UserId = jsonBody.UserId
	request.Genre = jsonBody.Genre
	request.Comment = jsonBody.Comment
	request.Image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, obj)

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
