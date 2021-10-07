package service

import (
	"Movie_Database/Repository"
	"Movie_Database/entities"

	"github.com/google/uuid"
)

type Service struct {
	Repo Repository.Repo
}

func DoService(r Repository.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) AddMovie(m entities.Movie) error {
	m.Id = uuid.New().String()

	err := s.Repo.AddMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetMovieById(id string) (entities.Movie, error) {
	movie, err := s.Repo.GetMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s Service) DeleteMovieById(id string) error {

	err := s.Repo.DeleteMovieById(id)
	if err != nil {
		return err
	}
	return nil
}

//func (s Service) UpdateMovieDb(mv entities.Movie) error{
//
//}
