package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./decks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS decks (
			id TEXT PRIMARY KEY,
			shuffled INTEGER,
			remaining INTEGER
		)
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}
