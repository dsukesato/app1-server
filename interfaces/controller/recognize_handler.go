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

type RecognizeController struct {
	Interactor usecase.RecognizeInteractor
}

func NewRecognizeController(dbconn database.DBConn) *RecognizeController {
	return &RecognizeController {
		Interactor: usecase.RecognizeInteractor {
			RecognizeRepository: &database.RecognizeRepository {
				DBConn: dbconn,
			},
		},
	}
}

func (c *RecognizeController) RecognizeIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/recognize/" {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	rec, err := c.Interactor.RecognizeAll(ctx)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(rec); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RecognizeController) RecognizeIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["id"])
	id, err := strconv.Atoi(params["id"])
	log.Printf("id: %d\n", id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	recognize, err := c.Interactor.RecognizeById(ctx, id)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(recognize); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RecognizeController) RecognizeUIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf(params["user_id"])
	uid, err := strconv.Atoi(params["user_id"])
	log.Printf("user_id: %d\n", uid)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ctx := r.Context()
	recognize, err := c.Interactor.RecognizeByUId(ctx, uid)

	if err != nil {
		log.Printf("err: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(recognize); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (c *RecognizeController) RecognizeSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/recognize/" {
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
	var jsonBody = new(model.PostRecognizeRequest)
	err = json.Unmarshal(body[:length], &jsonBody) // json -> Go Object
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := model.PostRecognizeRequest{}
	request.RestaurantId = jsonBody.RestaurantId
	request.UserId = jsonBody.UserId

	recognize, err := c.Interactor.Add(ctx, request)

	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(recognize); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
