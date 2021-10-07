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

func (r *Repo) DeleteMovieById(id string) error {
	db := MvDb{}                             // instance of slice of Movies
	file, err := ioutil.ReadFile(r.Filename) // reads moviedb.json file
	if err != nil {
		fmt.Println(err) // checks for errors
	}

	err = json.Unmarshal(file, &db) // convert to Go
	if err != nil {
		return err
	}

	for i, v := range db.Movies { // iterate through the movies struct via datatbase
		if v.Id == id { // compare and check if ID is in database
			db.Movies = append(db.Movies[:i], db.Movies[i+1:]...) // removes the movie selected from slice of movies and reappends the slice back together
		}
		mvBytes, err := json.MarshalIndent(&db, "", " ") // converts to Json
		if err != nil {
			return err
		}
		_ = ioutil.WriteFile(r.Filename, mvBytes, 0644) // writes the change bak to the db file.
		return nil
	}
	return nil
}

//func (r *Repo) UpdateMovieDb (mv entities.Movie) error {
//
//}
