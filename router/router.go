package router

import (

	"github.com/gorilla/mux"
	"github.com/mongo-golang-hitesh/controller"
)

func MyRouter() *mux.Router{

	r:=mux.NewRouter()

	r.HandleFunc("/movie",controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/deleteall",controller.DeleteAllMovie).Methods("DELETE")
return r
}