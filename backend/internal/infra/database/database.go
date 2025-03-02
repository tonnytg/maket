package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDatabase(conn *sql.DB) {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS targets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			value REAL,
			deadline TEXT,
			created_at TEXT
		);
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *sql.DB {

	pool := &sql.DB{}
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)
	pool.SetConnMaxLifetime(0)

	pool, err := sql.Open("sqlite3", "./makemoneytarget.db")
	if err != nil {
		log.Fatal(err)
	}

	return pool
}

func Start() {

	conn := GetConnection()
	defer conn.Close()

	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	InitDatabase(conn)
}
