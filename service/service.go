package service

import (
	//"Movie_Database/Repository"
	//"Movie_Database/Repository"
	"Movie_Database/entities"
	"github.com/google/uuid"
	"net/http"
)

type Repository interface {
	AddMovie(m entities.Movie) error
	GetMovieById(id string) (entities.Movie, error)
	DeleteMovieById(id string) error
	UpdateMovieById(id string, mv entities.Movie) error
}

type Serv struct {
	Repo Repository
}

func DoService(r Repository) Serv {
	return Serv{
		Repo: r,
	}
}

func (s *Serv) AddMovie(m entities.Movie) error {
	m.Id = uuid.New().String()

	err := s.Repo.AddMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Serv) GetMovieById(id string) (entities.Movie, error) {
	movie, err := s.Repo.GetMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s *Serv) DeleteMovieById(id string) error {

	err := s.Repo.DeleteMovieById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Serv) UpdateMovieById(id string, mv entities.Movie) error {
	if id != mv.Id {
		return http.ErrMissingFile
	}
	err := s.Repo.UpdateMovieById(id, mv)
	if err != nil {
		return err
	}
	return nil
}
