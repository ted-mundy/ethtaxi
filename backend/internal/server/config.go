package server

import (
	"log/slog"
	"os"
	"strconv"
)

type ServerConfig struct {
	Port int `json:"port"`
}

func getServerConfig() *ServerConfig {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		slog.Debug("Failed to parse SERVER_PORT environment variable, using default port", "error", err)
		port = 8080 // default port
	}

	return &ServerConfig{
		Port: port,
	}
}
