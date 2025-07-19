package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/ted-mundy/ethtaxi/internal/server"
)

func main() {
	// not going to follow the cmd/ convention, just a simple main function
	// as this will only be used for hosting the backend API
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Failed to load .env file, proceeding without it", "error", err)
	}
	s, err := server.NewServer()

	if err != nil {
		panic("failed to create server: " + err.Error())
	}

	slog.Info("Starting backend server...")

	if err := s.Start(); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
