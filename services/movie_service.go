package services

import (
	"Stockbit/models"
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

type MovieService interface {
	AllMovies (r *http.Request) *models.Result
	DetailMovie (w http.ResponseWriter, r *http.Request)
}

type movieService struct {
	repoMovie repositories.AllMoviesRepository
	repoDB repositories.LogRepository
	repoMov repositories.MovieRepository
}

func NewMovieService (a repositories.AllMoviesRepository, b repositories.LogRepository) *movieService {
	return &movieService{
		repoMovie: a,
		repoDB: b,
	}
}

func (ms movieService) AllMovies (r *http.Request) *models.Result {
	res, err, url := ms.repoMovie.GetMovies(r)
	if err != nil {
		log.Fatal(err)
	}

	err = ms.repoDB.Store(url)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (ms movieService) DetailMovie (w http.ResponseWriter, r *http.Request)  {
	Movie, err, url := ms.repoMov.GetMovie(r)
	if err != nil {
		log.Println(err)
	}
	err = ms.repoDB.Store(url)
	if err != nil {
		log.Println(err)
	}
	_ = json.NewEncoder(w).Encode(Movie)

}
