package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func OpenConnection() (*sql.DB, error) {
	connStr := "user=postgres dbname=postgres password=1212 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println("error closing database connection: ", err)
	}
}

func InsertExcuse(db *sql.DB, text string) error {
	_, err := db.Exec("INSERT INTO excuses(string, count) VALUES ($1, $2)", text, 0)
	return err
}
