package main

import (
	"Stockbit/services"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"net/http"
)

func main() {
	mux := mux2.NewRouter()

	mux.HandleFunc("/allMovies", services.AllMovies).Methods("GET")
	mux.HandleFunc("/movie", services.DetailMovie).Methods("GET")

	var handler http.Handler = mux

	//serve a server
	server := new(http.Server)
	server.Addr = ":8000"
	server.Handler = handler

	fmt.Println("server started at localhost:8000")
	_ = server.ListenAndServe()
}
