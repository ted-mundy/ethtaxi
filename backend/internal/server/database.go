package server

import (
	"database/sql"
	"fmt"
)

func checkDatabaseSetup(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// could probably check schema but not going to do that now

	return nil
}
