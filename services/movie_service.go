package services

import (
	"Stockbit/repositories"
	"encoding/json"
	"log"
	"net/http"
)
	/*Developer Note's:
		I know this might be look weird, I made a interface contract
		in repository but didn't use on this service. I should make delivery service
		to handle the http request for this service instead of doing this, but the problem is
		what i learn for clean architecture, most of it use context as param and communication
		but i still didn't know how to use context properly...*/

//type AllMoviesService struct {
//	Movies allMovies.AllMoviesRepository
//}

func AllMovies (w http.ResponseWriter, r *http.Request)  {
	Movies, err, url := repositories.GetMovies(r)
	if err != nil {
		log.Println(err)
	}
	err = repositories.Store(url)
	if err != nil {
		log.Println(err)
	}
	_ = json.NewEncoder(w).Encode(Movies)
}

func DetailMovie (w http.ResponseWriter, r *http.Request)  {
	Movie, err, url := repositories.GetMovie(r)
	if err != nil {
		log.Println(err)
	}
	err = repositories.Store(url)
	if err != nil {
		log.Println(err)
	}
	_ = json.NewEncoder(w).Encode(Movie)

}
