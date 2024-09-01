package db

import (
	"database/sql"
	"log" // Use log for error logging instead of panic

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to the global DB variable
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5) // At least 5 connections open at all times

	createTables()
}

func createTables() {
	if DB == nil {
		log.Fatal("DB is not initialized") // Additional safety check
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create events table:", err)
	}
}
