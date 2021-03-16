package delivery

import (
	"Stockbit/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type movieDelivery struct {
	service services.MovieService
}

func NewMoviesHandler(service services.MovieService, mx *mux.Router) {
	handler := &movieDelivery{
		service: service,
	}

	mx.HandleFunc("/allMovies", handler.GetMoviesDelivery).Methods("GET")
}

func (mv movieDelivery) GetMoviesDelivery (w http.ResponseWriter, r *http.Request) {
	result := mv.service.AllMovies(r)
	_ = json.NewEncoder(w).Encode(result)
}
