package Repository

import (
	"encoding/json"
	"io/ioutil"
)

func CovertFileData() (MvDb, error) {
	////create variable
	data := MvDb{}
	////first read the file							//absolute path
	jsonBytes, err := ioutil.ReadFile("/home/heather/Assignment!/Movie_Database/moviedb.json")
	if err != nil {
		return data, err
	}
	//second
	err = json.Unmarshal(jsonBytes, &data) //all the data and memory use the Unmarshal to grab all as one piece
	if err != nil {
		return data, err
	}
	return data, nil

}

func ConvertDataBack(db MvDb) error {
	movieBytes, err := json.Marshal(db)
	if err != nil {
		return err
	}
	//this function overwrites the file by default
	err = ioutil.WriteFile("/home/heather/Assignment!/Movie_Database/moviedb.json", movieBytes, 0644)
	return nil
}
