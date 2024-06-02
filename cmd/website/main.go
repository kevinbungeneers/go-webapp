package main

import (
	"go-webapp/internal/server"
	"log/slog"
)

func main() {
	s := server.NewWebServer(
		server.WithListenAddress("localhost:3000"),
		server.WithLogLevel(slog.LevelDebug),
	)

	s.Start()
}
