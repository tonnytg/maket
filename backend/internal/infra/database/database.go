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
		    uuid TEXT NOT NULL,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    status TEXT NOT NULL,
		    target_amount REAL NOT NULL,
		    current_amount REAL NOT NULL,
		    start_date TEXT NOT NULL,
		    end_date TEXT NOT NULL
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS target_members (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    uuid TEXT NOT NULL,
		    name TEXT NOT NULL,
		    email TEXT NOT NULL,
			role TEXT NOT NULL,
		    target_id INTEGER NOT NULL,
		    FOREIGN KEY(target_id) REFERENCES targets(id)
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
