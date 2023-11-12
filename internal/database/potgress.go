package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
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

func Init() error {
	newdataBase := GetDB()
	if err := createAccountTable(newdataBase); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createAccountTable(s *sql.DB) error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`
	_, err := s.Exec(query)
	return err
}
