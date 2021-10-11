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

	file, err := ioutil.ReadFile(r.Filename) // reads the file moviedb.json
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &db) // pointer indicates change and points to address of the instance of the Movie Struct
	if err != nil {
		return err
	}

	db.Movies = append(db.Movies, m) //appends slice of movies to the instance of mvDB struct

	movieBytes, err := json.MarshalIndent(db, "", "	") // takes in the bytes and converts back to json
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644) //writes the information to the database file
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) GetMovieById(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename) // reads the moviedb.json file
	if err != nil {
		fmt.Println(err)
	}
	movies := MvDb{} // instance of slice of movies
	err = json.Unmarshal(file, &movies)

	compare := entities.Movie{} // instance of the Movie struct

	for _, val := range movies.Movies { // iterates through the slice of movies and the struct setup
		if val.Id == id { // compares the id entered in postman to the movie db id stored.
			compare = val
			return compare, nil // returns the movie found by id or error if not found

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

	newMovieDb := MvDb{}

	for i, v := range db.Movies {
		if v.Id == id {
			newMovieDb.Movies = append(db.Movies[:i], db.Movies[i+1:]...)
		}

	}
	mvBytes, err := json.MarshalIndent(newMovieDb, "", " ") // converts to Json
	if err != nil {
		return err
	}
	_ = ioutil.WriteFile(r.Filename, mvBytes, 0644) // writes the change bak to the db file.
	if err != nil {
		return err
	}
	return nil

}

func (r *Repo) UpdateMovieDb(id string, mv entities.Movie) error {

	file, err := ioutil.ReadFile(r.Filename) // reads moviedb.json file
	if err != nil {
		return err
	}

	db := MvDb{}
	err = json.Unmarshal(file, &db) // convert to Go
	if err != nil {
		return err
	}

	for i, v := range db.Movies {
		if v.Id == id {
			db.Movies[i] = mv
		}

	}

	mvBytes, err := json.MarshalIndent(db.Movies, "", " ") // converts to Json
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, mvBytes, 0644) // writes the change bak to the db file.
	if err != nil {
		return err
	}
	return nil
}
