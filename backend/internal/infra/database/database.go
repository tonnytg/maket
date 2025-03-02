package database

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"time"
)

func InitDatabase() {
}

var pool *sql.DB // Database connection pool.

func Start() {

	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")

	pool, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(); err != nil {
		panic(err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

}

func Query(query string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

}
