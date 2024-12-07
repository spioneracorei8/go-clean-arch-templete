package main

import (
	"go-clean-arch-templete/config"
	"go-clean-arch-templete/server"
)

func getMainServer() *server.Server {
	return &server.Server{
		APP_PORT:          config.APP_PORT,
		APP_LOGGER:        config.APP_LOGGER,
		// // cors
		ALLOW_HEADERS:     config.ALLOW_HEADERS,
		ALLOW_ORIGINS:     config.ALLOW_ORIGINS,
		ALLOW_METHODS:     config.ALLOW_METHODS,
		ALLOW_CREDENTIALS: config.ALLOW_CREDENTIALS,
		// // mongo db
		// MONGODB_CONNECTION_URI: config.MONGODB_CONNECTION_URI,
	}
}

func main() {
	server := getMainServer()
	server.Start()
}
