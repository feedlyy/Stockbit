package repositories

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet (t *testing.T) {

	t.Run("Movies", func(t *testing.T) {
		result, _, url := GetMovies(httptest.NewRequest("GET",
			"http://www.omdbapi.com/?apikey=faf7e5bb&&page=3&s=Godzilla", nil))

		apikey := strings.Index(url, "apikey")
		apikeys := url[apikey+7:39]
		assert.NotNil(t, url)
		assert.NotNil(t, result)
		assert.Equal(t, "faf7e5bb", apikeys)
	})

	t.Run("Movie", func(t *testing.T) {
		result, _, url := GetMovie(httptest.NewRequest("GET",
			"http://www.omdbapi.com/?apikey=faf7e5bb&t=Godzilla", nil))

		apikey := strings.Index(url, "apikey")
		apikeys := url[apikey+7:39]
		assert.NotNil(t, url)
		assert.NotNil(t, result)
		assert.Equal(t, "faf7e5bb", apikeys)
	})
}