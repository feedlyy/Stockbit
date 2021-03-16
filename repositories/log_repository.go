package repositories

import (
	"database/sql"
	"time"
)

type LogRepository interface {
	Store(url string) error
}

type logRepo struct {
	DB *sql.DB
}

func NewLogRepo (db *sql.DB) *logRepo {
	return &logRepo{
		DB: db,
	}
}

func (lr logRepo) Store (url string) error {
	query := "insert into logs(url, created_at) values (?, ?)"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	_, err := lr.DB.Exec(query, url, time.Now().In(loc))
	if err != nil {
		return err
	}
	return nil
}
