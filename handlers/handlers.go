package handlers

import (
	"Movie_Database/entities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Service interface {
	AddMovie(m entities.Movie) error
	GetMovieById(id string) (entities.Movie, error)
	DeleteMovieById(id string) error
	UpdateMovieById(id string, mv entities.Movie) error
}

type MovieHandler struct {
	Serv Service
}

func NewMovieHandler(s Service) MovieHandler {
	return MovieHandler{
		Serv: s,
	}
}

func (mov MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = mov.Serv.AddMovie(mv)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (mov MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	movById, err := mov.Serv.GetMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movie, err := json.MarshalIndent(movById, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(movie)
}

func (mov MovieHandler) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mov.Serv.DeleteMovieById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (mov MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = mov.Serv.UpdateMovieById(id, mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
