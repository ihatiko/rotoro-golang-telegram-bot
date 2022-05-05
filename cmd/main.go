package cmd

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
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
	router := httprouter.New()

	router.POST("/", handlers.Handler)
	router.GET("/health", handlers.HealthCheckHandler)
	log.Info().Msgf("start server on port %s", s.Port)
	err := http.ListenAndServe(s.Port, router)
	if err != nil {
		log.Fatal().Err(err).Msg("http.ListenAndServe")
	}
}
