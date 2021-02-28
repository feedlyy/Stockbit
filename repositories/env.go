package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/stockbit")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
