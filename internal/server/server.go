package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	HttpServer *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleIndex(w, r, logger)
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleUpload(w, r, logger)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{logger: logger, HttpServer: server}
}
