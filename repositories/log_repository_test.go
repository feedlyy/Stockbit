package repositories

import (
	"Stockbit/models"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
// this for mocking test time.Time
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestLogRepo_Store(t *testing.T) {
	ar := &models.Logs{
		Id: 1,
		Urls:       "http://www.omdbapi.com/?apikey=faf7e5bb&&page=3&s=Godzilla",
		Created_at: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "insert into logs\\(url, created_at\\)"
	mock.ExpectExec(query).WithArgs(ar.Urls, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1,1))

	log := NewLogRepo(db)
	err = log.Store(ar.Urls)
	assert.NoError(t, err)
}
