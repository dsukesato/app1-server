package controller

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	usecase "github.com/dsukesato/go13/pbl/app1-server/usecase/interactor"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type RestaurantsController struct {
	Interactor usecase.RestaurantsInteractor
}

func NewRestaurantsController(dbconn database.DBConn) *RestaurantsController {
	return &RestaurantsController {
		Interactor: usecase.RestaurantsInteractor {
			RestaurantsRepository: &database.RestaurantsRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *RestaurantsController) RestaurantsIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	rests, err := c.Interactor.RestaurantsAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(rests); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RestaurantsController) RestaurantsIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["id"])
	id, err := strconv.Atoi(params["id"])
	log.Printf("id: %d\n", id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	rest, err := c.Interactor.RestaurantById(ctx, id)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(rest); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RestaurantsController) RestaurantsSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
		http.NotFound(w, r)
		return
	}

	//if r.Header.Get("Content-Type") != "multipart/form-data" {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	formValue := r.FormValue("json")

	var jsonBody model.PostRestaurantRequest

	b := []byte(formValue)
	err := json.Unmarshal(b, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	formFile, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	defer formFile.Close()

	ctx := r.Context()
	rLastId, err := c.Interactor.RestaurantLastId(ctx)

	bucket := "pbl-lookin-storage" // GCSバケット名
	obj := fmt.Sprintf("restaurants/restaurant%d.jpeg", rLastId+1)
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
	if err != nil {
		log.Printf("err: %v\n", err)
	}

	if err := writer.Close(); err != nil {
		log.Printf("failed to close gcs writer : %v", err)
	}

	request := model.RestaurantRequest{}
	request.Uuid = jsonBody.Uuid
	request.Name = jsonBody.Name
	request.BusinessHours = jsonBody.BusinessHours
	request.Image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, obj)

	posts, err := c.Interactor.Add(ctx, request)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RestaurantsController) RestaurantsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
		http.NotFound(w, r)
		return
	}

	//chk := true

	formValue := r.FormValue("json")

	var jsonBody model.PutRestaurantRequest

	b := []byte(formValue)
	err := json.Unmarshal(b, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	request := model.PutRestaurantRequest{}
	request.Id = jsonBody.Id
	request.Uuid = jsonBody.Uuid
	request.Name = jsonBody.Name
	request.BusinessHours = jsonBody.BusinessHours

	formFile, _, err := r.FormFile("image")
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	defer formFile.Close()

	ctx := r.Context()

	bucket := "pbl-lookin-storage" // GCSバケット名
	obj := fmt.Sprintf("restaurants/restaurant%d.jpeg", request.Id)
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
	if err != nil {
		log.Printf("err in io.Copy: %v\n", err)
	}

	if err := writer.Close(); err != nil {
		log.Printf("failed to close gcs writer : %v", err)
	}

	request.Image = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, obj)

	restaurant, err := c.Interactor.Update(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(restaurant); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
