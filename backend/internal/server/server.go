package server

import (
	"fmt"
	"net/http"
)

func Start() error {
	config := getServerConfig()

	return http.ListenAndServe(":"+fmt.Sprintf("%d", config.Port), nil)
}
