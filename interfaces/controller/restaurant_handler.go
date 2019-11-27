package controller

import (
	"encoding/json"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	usecase "github.com/dsukesato/go13/pbl/app1-server/usecase/interactor"
)

type RestaurantsController struct {
	Interactor usecase.RestaurantsInteractor
}

func NewRestaurantsController(dbconn database.DBConn) *RestaurantsController {
	return &RestaurantsController {
		Interactor: usecase.RestaurantsInteractor{
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

type GetRestaurantsResponse struct {
	Restaurants []RestaurantsField `json:"restaurants"`
}

type RestaurantsField struct {
	RestaurantId int       `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at"`
}

func (c *RestaurantsController) RestaurantsSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
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
	var jsonBody = new(model.PostRestaurantRequest)
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostRestaurantRequest{}
	request.Name = jsonBody.Name
	request.BusinessHours = jsonBody.BusinessHours
	request.Image = jsonBody.Image

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

type PostRestaurantResponse struct {
	RestaurantId int       `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
	Message string         `json:"massage"`
}
