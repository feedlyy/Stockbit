package main

import (
	"Stockbit/delivery"
	"Stockbit/repositories"
	"Stockbit/services"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	db, err := repositories.Connect()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redis := repositories.ConnectRedis(10)

	mux := mux2.NewRouter()
	repoMovie := repositories.NewAllMovies(redis)
	repoLogDB := repositories.NewLogRepo(db)
	serviceMovie := services.NewMovieService(repoMovie, repoLogDB)
	delivery.NewMoviesHandler(serviceMovie, mux)

	var handler http.Handler = mux

	//serve a server
	server := new(http.Server)
	server.Addr = ":8000"
	server.Handler = handler

	fmt.Println("server started at localhost:8000")
	_ = server.ListenAndServe()
}
