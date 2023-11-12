package model

import (
	"database/sql"
)

type PostgresStore struct {
	DB *sql.DB
}