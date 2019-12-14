package controller

import (
	"encoding/json"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	usecase "github.com/dsukesato/go13/pbl/app1-server/usecase/interactor"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type PointController struct {
	Interactor usecase.PointInteractor
}

func NewPointController(dbconn database.DBConn) *PointController {
	return &PointController {
		Interactor: usecase.PointInteractor {
			PointRepository: &database.PointRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *PointController) PointIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/point/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	points, err := c.Interactor.PointAll(ctx)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(points); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *PointController) PointIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["id"])
	id, err := strconv.Atoi(params["id"])
	log.Printf("id: %d\n", id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	point, err := c.Interactor.PointById(ctx, id)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(point); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *PointController) PointSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/point/" {
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
	var jsonBody = new(model.PostPointRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostPointRequest{}
	request.RestaurantId = jsonBody.RestaurantId
	request.UserId = jsonBody.UserId
	request.Transaction = jsonBody.Transaction

	point, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(point); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
