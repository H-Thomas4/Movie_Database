package handlers

import (
	"github.com/gorilla/mux"
)

func NewServer(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie/{Id}", handler.GetById).Methods("GET")
	r.HandleFunc("/movie/delete/{Id}", handler.DeleteMovieById).Methods("DELETE")

	return r
}
