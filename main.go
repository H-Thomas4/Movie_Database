package main

import (
	"Movie_Database/Repository"
	"Movie_Database/handlers"
	"Movie_Database/service"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fileName := "/home/heather/Assignment!/Movie_Database/moviedb.json"
	ext := filepath.Ext(fileName)

	if ext != ".json" {
		log.Fatalln("incorrect file extension")
	}

	repository := Repository.NewRepo(fileName)
	serv := service.DoService(repository)
	handler := handlers.NewMovieHandler(serv)
	router := handlers.NewServer(handler)

	//repository2 := Repository.NewRepo(fileName)
	//serv2 := service.DoService(repository2)
	//handler2 := handlers.NewMovieHandler(serv2)
	//router2 := handlers.NewServer(handler2)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	fmt.Println("Successfully running server 8080")

	log.Fatal(server.ListenAndServe())
}
