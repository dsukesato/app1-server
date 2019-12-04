package infrastructure

import (
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Serve() {
	r := mux.NewRouter()
	// 依存関係注入
	// restaurants(飲食店)
	rc := controller.NewRestaurantsController(Init())
	// posts(投稿)
	pc := controller.NewPostsController(Init()) // pcはpostsControllerの略
	// users(ユーザ)
	uc := controller.NewUsersController(Init())
	// recognize(AR認識)
	rec := controller.NewRecognizeController(Init())
	// good(いいね)
	gc := controller.NewGoodController(Init())

	r.HandleFunc("/Lookin/restaurants/", rc.RestaurantsIndexHandler).Methods("GET")
	r.HandleFunc("/Lookin/restaurants/{id}", rc.RestaurantsIdHandler).Methods("GET")
	r.HandleFunc("/Lookin/restaurants/", rc.RestaurantsSendHandler).Methods("POST")

	r.HandleFunc("/Lookin/posts/", pc.PostsIndexHandler).Methods("GET")
	r.HandleFunc("/Lookin/posts/{id}", pc.PostsIdHandler).Methods("GET")
	// RI=RestaurantId, G=Genre
	r.HandleFunc("/Lookin/posts/restaurant_id:{restaurant_id}/genre:{genre}", pc.PostsRIGHandler).Methods("GET")
	r.HandleFunc("/Lookin/posts/", pc.PostsSendHandler).Methods("POST")

	r.HandleFunc("/Lookin/users/", uc.UsersIndexHandler).Methods("GET")
	r.HandleFunc("/Lookin/users/{id}", uc.UsersIdHandler).Methods("GET")
	r.HandleFunc("/Lookin/users/", uc.UsersSendHandler).Methods("POST")

	r.HandleFunc("/Lookin/recognize/", rec.RecognizeIndexHandler).Methods("GET")
	r.HandleFunc("/Lookin/recognize/id:{id}", rec.RecognizeIdHandler).Methods("GET")
	r.HandleFunc("/Lookin/recognize/user_id:{user_id}", rec.RecognizeUIdHandler).Methods("GET")
	r.HandleFunc("/Lookin/recognize/", rec.RecognizeSendHandler).Methods("POST")

	r.HandleFunc("/Lookin/good/", gc.GoodIndexHandler).Methods("GET")
	r.HandleFunc("/Lookin/good/{id}", gc.GoodIdHandler).Methods("GET")
	r.HandleFunc("/Lookin/good/", gc.GoodSendHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
