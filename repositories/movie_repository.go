package repositories

import (
	"Stockbit/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MovieRepository interface {
	GetMovie() (*models.Movie, error, string)
}

func GetMovie(r *http.Request) (*models.Movie, error, string) {
	var err error
	var client = &http.Client{}
	var film models.Movie

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err, ""
	}
	request.Header.Set("Content-Type", "application/json")

	q := request.URL.Query()
	q.Add("apikey", r.URL.Query()["apikey"][0])
	q.Add("t", r.URL.Query()["t"][0])
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
		err = json.Unmarshal(bodyBytes, &film)
		if err != nil {
			return nil, nil, ""
		}
	}
	return &film, nil, request.URL.String()
}
