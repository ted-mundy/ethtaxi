package database

import (
	"fmt"
	"os"
)

type databaseConfig struct {
	connectionString string
}

func getDatabaseConfig() (*databaseConfig, error) {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	return &databaseConfig{
		connectionString: connectionString,
	}, nil
}
