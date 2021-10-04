package Repository

import (
	"Movie_Database/entities"
	"encoding/json"
	"fmt"

	//"errors"
	"io/ioutil"
)

type MvDb struct {
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo {
	return Repo{
		Filename: f,
	}
}

func (r Repo) AddMovie(m entities.Movie) error {
	db := MvDb{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &db)
	if err != nil {
		return err
	}

	db.Movies = append(db.Movies, m)

	movieBytes, err := json.MarshalIndent(db, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) GetMovieById(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}
	movies := MvDb{}
	err = json.Unmarshal(file, &movies)

	compare := entities.Movie{}

	for _, val := range movies.Movies {
		if val.Id == id {
			compare = val
			return compare, nil

		}

	}
	return compare, nil
}
