package repositories

import (
	"time"
)

type LogRepository interface {
	Store() error
}

func Store (url string) error {
	db := Connect()

	query := "insert into logs(url, created_at) values (?, ?)"

	loc, _ := time.LoadLocation("Asia/Jakarta")

	_, err := db.Exec(query, url, time.Now().In(loc))
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}
