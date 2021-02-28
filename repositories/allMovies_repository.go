package repositories

import (
	"Stockbit/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var url = "http://www.omdbapi.com/"

type AllMoviesRepository interface {
	GetMovies() (*models.Result, error, string)
}

func GetMovies(r *http.Request) (*models.Result, error, string) {
	var err error
	var client = &http.Client{}
	var films models.Result

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err, ""
	}
	request.Header.Set("Content-Type", "application/json")

	q := request.URL.Query()
	q.Add("apikey", r.URL.Query()["apikey"][0])
	q.Add("page", r.URL.Query()["page"][0])
	q.Add("s", r.URL.Query()["s"][0])
	request.URL.RawQuery = q.Encode()

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
		err = json.Unmarshal(bodyBytes, &films)
		if err != nil {
			return nil, nil, ""
		}
	}
	return &films, nil, request.URL.String()
}


