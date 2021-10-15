package handlers

import (
	"github.com/gorilla/mux"
)

//type MovieHandler interface {
//	PostNewMovie(w http.ResponseWriter, r *http.Request)
//	GetById(w http.ResponseWriter, r *http.Request)
//	DeleteMovieById(w http.ResponseWriter, r *http.Request)
//	UpdateMovie(w http.ResponseWriter, r *http.Request)
//}

func NewServer(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie/{Id}", handler.GetById).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.DeleteMovieById).Methods("DELETE")
	r.HandleFunc("/movie/{Id}", handler.UpdateMovie).Methods("PUT")

	return r
}
