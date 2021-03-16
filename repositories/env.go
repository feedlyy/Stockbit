package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/stockbit")

	if err != nil {
		return nil, err
	}

	return db, nil
}

var pool *redis.Pool

func ConnectRedis(maxConnection int) *redis.Pool {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: maxConnection,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}

	return pool
}