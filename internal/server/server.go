package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type server struct {
	listenAddress string
	logLevel      slog.Level
}

func (s server) Start(handler http.Handler) {
	ctx := context.Background()

	configureLogger(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     s.logLevel,
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	httpServer := http.Server{
		Addr:    s.listenAddress,
		Handler: handler,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-stop

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func newServer(opts ...Option) server {
	s := server{
		listenAddress: "localhost:8080",
		logLevel:      slog.LevelInfo,
	}

	for i := range opts {
		opts[i](&s)
	}

	return s
}
