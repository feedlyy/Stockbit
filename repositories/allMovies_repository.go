package repositories

import (
	"Stockbit/models"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var url = "http://www.omdbapi.com/"

type AllMoviesRepository interface {
	GetMovies(r *http.Request) (*models.Result, error, string)
}

type moviesRepository struct {
	rediss *redis.Pool
	films *models.Result
}

func NewAllMovies (redis *redis.Pool) *moviesRepository {
	return &moviesRepository{rediss: redis}
}

func (movies moviesRepository) GetMovies(r *http.Request) (*models.Result, error, string) {
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err, ""
	}

	title := r.URL.Query()["s"][0]
	page := r.URL.Query()["page"][0]
	apikey := r.URL.Query()["apikey"][0]

	q := request.URL.Query()
	q.Add("apikey", apikey)
	q.Add("page", page)
	q.Add("s", title)
	request.URL.RawQuery = q.Encode()

	conn := movies.rediss.Get()
	defer conn.Close()

	//check if req already cached
	cached, err := redis.Values(conn.Do("GET", strings.ToLower(title) + "-" + page))
	if err == nil {
		err = redis.ScanStruct(cached, &movies.films)
		if err != nil {
			log.Fatal(err)
		}
		return movies.films, nil, request.URL.String()
	}

	//do request if didn't
	response, err := client.Do(request)
	if err != nil{
		return nil, err, ""
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bodyBytes, &movies.films)
		if err != nil {
			return nil, err, ""
		}

		// simpan data response di redis
		_, err = conn.Do("SET", strings.ToLower(title) + "-" + page, string(bodyBytes))
		if err != nil {
			log.Fatal(err)
		}
	}
	return movies.films, nil, request.URL.String()
}


