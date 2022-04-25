package cmd

import (
	"net/http"
	"os"
	"rotoro-golang-telegram-bot/internal"
)

type Server struct {
	Port string
}

func NewServer() *Server {
	port := os.Getenv(port)
	if port == empty {
		port = defaultPort
	}
	return &Server{Port: port}
}

func (s *Server) Serve() {
	handlers := internal.NewHandlers()
	http.ListenAndServe(s.Port, http.HandlerFunc(handlers.Handler))
}
