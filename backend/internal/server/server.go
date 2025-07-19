package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ted-mundy/ethtaxi/internal/database"
)

type Server struct {
	db     *sql.DB
	router *http.ServeMux
	config *ServerConfig
}

func NewServer() (*Server, error) {
	router := http.NewServeMux()

	config := getServerConfig()
	db, err := database.GetDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	return &Server{
		db,
		router,
		config,
	}, nil
}

func (s *Server) Start() error {
	err := checkDatabaseSetup(s.db)
	if err != nil {
		return fmt.Errorf("failed to check database setup: %w", err)
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router)
}
