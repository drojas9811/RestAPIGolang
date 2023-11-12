package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// DatabaseURL represents the connection string to the database
const DatabaseURL = "host=db port=5432 user=postgres dbname=postgres password=gobank sslmode=disable"

var (
	db   *sql.DB
	once sync.Once
)

// GetDB returns a singleton instance of the database connection
func GetDB() *sql.DB {
	once.Do(func() {
		// Initialize the database connection
		conn, err := sql.Open("postgres", DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		db = conn

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to the database")
	})

	return db
}