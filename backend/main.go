package main

import (
	"log/slog"

	"github.com/ted-mundy/ethtaxi/internal/server"
)

func main() {
	// not going to follow the cmd/ convention, just a simple main function
	// as this will only be used for hosting the backend API
	slog.Info("Starting backend server...")
	if err := server.Start(); err != nil {
		slog.Error("Failed to start server - panicking!")
		panic("failed to start server: " + err.Error())
	}
}
