package server

import (
	"go-webapp/internal/http/adapter"
	"go-webapp/internal/http/middleware"
	"go-webapp/internal/template"
	"go-webapp/web"
	"log"
	"net/http"
)

type WebServer struct {
	server  server
	handler http.Handler
}

func (w WebServer) Start() {
	w.server.Start(w.handler)
}

func NewWebServer(opts ...Option) WebServer {
	r, err := template.NewRenderer()
	if err != nil {
		log.Fatal(err)
	}

	publicFS, err := web.PublicFS()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /{$}", adapter.NewHomePageAdapter(r))
	mux.Handle("GET /", http.FileServer(http.FS(publicFS)))

	handler := middleware.Logger(mux)
	handler = middleware.RequestID(handler)

	return WebServer{
		handler: handler,
		server:  newServer(opts...),
	}
}
