package server

type ServerConfig struct {
	Port int `json:"port"`
}

func getServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: 8080,
	}
}
