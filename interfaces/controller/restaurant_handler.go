package controller

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	usecase "github.com/dsukesato/go13/pbl/app1-server/usecase/interactor"
	"github.com/gorilla/mux"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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

	flag.BoolVar(&inMemory, "in-mem", false, "Add -in-mem flag for in-memory-only uploads")
	flag.Parse()

	formValue := r.FormValue("json")

	var jsonBody model.PostRestaurantRequest

	b := []byte(formValue)
	err := json.Unmarshal(b, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	formFile, _, err := r.FormFile("image")
	handleError(err)
	defer formFile.Close()

	dir, err := os.Getwd()
	handleError(err)

	filename := "upload_restaurant.jpeg"
	saveFile, err := os.Create(path.Join(dir + "/image", filename))
	handleError(err)
	defer saveFile.Close()

	if inMemory {
		_, err = io.Copy(saveFile, formFile)
	} else {
		uploadFile, err := os.Create(path.Join(dir + "/image", filename))
		handleError(err)
		defer uploadFile.Close()

		_, err = io.Copy(uploadFile, formFile)
	}
	handleError(err)

	//uploadFile, err := os.Create(path.Join(dir + "/image", filename))
	//handleError(err)
	//_, err = io.Copy(uploadFile, formFile)

	// gcs
	file, err := os.Open("image/upload_restaurant.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	image, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}
	imageBytes := buffer.Bytes()

	ctx := r.Context()
	id, err := c.Interactor.RestaurantLastId(ctx)

	bucket := "pbl-lookin-storage" // GCSバケット名
	obj := fmt.Sprintf("restaurant%d.jpeg", id+1)
	bCtx := context.Background()

	client, err := storage.NewClient(bCtx)
	if err != nil {
		log.Printf("failed to create gcs client : %v", err)
	}

	// GCS writer
	writer := client.Bucket(bucket).Object(obj).NewWriter(bCtx)
	writer.ContentType = "image/jpeg" // 任意のContentTypeに置き換える

	// upload : write object body
	if _, err := writer.Write(imageBytes); err != nil {
		log.Printf("failed to write object body : %v", err)
	}

	if err := writer.Close(); err != nil {
		log.Printf("failed to close gcs writer : %v", err)
	}
	w.WriteHeader(http.StatusCreated)

	request := model.RestaurantRequest{}
	request.Name = jsonBody.Name
	request.BusinessHours = jsonBody.BusinessHours
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
