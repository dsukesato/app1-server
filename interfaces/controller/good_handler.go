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

type GoodController struct {
	Interactor usecase.GoodInteractor
}

func NewGoodController(dbconn database.DBConn) *GoodController {
	return &GoodController {
		Interactor: usecase.GoodInteractor {
			GoodRepository: &database.GoodRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *GoodController) GoodIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/good/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	goods, err := c.Interactor.GoodAll(ctx)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(goods); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *GoodController) GoodIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["id"])
	id, err := strconv.Atoi(params["id"])
	log.Printf("id: %d\n", id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	good, err := c.Interactor.GoodById(ctx, id)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(good); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *GoodController) GoodSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/good/" {
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
	var jsonBody = new(model.PostGoodRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostGoodRequest{}
	request.PostId = jsonBody.PostId
	request.UserId = jsonBody.UserId
	//request.State = jsonBody.State

	good, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(good); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *GoodController) GoodUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/good/" {
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
	var jsonBody = new(model.PutGoodRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PutGoodRequest{}
	request.PostId = jsonBody.PostId
	request.UserId = jsonBody.UserId
	request.State = jsonBody.State

	good, err := c.Interactor.PutGood(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(good); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
