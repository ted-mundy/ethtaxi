package database

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func GetDatabaseConnection() (*sql.DB, error) {
	config, err := getDatabaseConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", config.connectionString)
	if err != nil {
		return nil, err
	}

	// we ping the database to ensure the connection is valid
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
