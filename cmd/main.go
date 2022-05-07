package cmd

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"rotoro-golang-telegram-bot/internal"
)

type Server struct {
	Port                 string
	TelegramRegistryHost string
}

func NewServer() *Server {
	port := os.Getenv(port)
	telegramHostRegistry := os.Getenv(telegramRegistryHost)
	if port == empty {
		port = defaultPort
	}
	return &Server{Port: port, TelegramRegistryHost: telegramHostRegistry}
}

func (s *Server) Serve() {
	handlers := internal.NewHandlers()

	router := s.GetRouter(handlers)
	log.Info().Msgf("start server on port %s", s.Port)
	err := http.ListenAndServe(s.Port, router)
	if err != nil {
		log.Fatal().Err(err).Msg("http.ListenAndServe")
	}
}

func (s *Server) GetRouter(handlers *internal.Handlers) *httprouter.Router {
	router := httprouter.New()
	router.POST("/", handlers.Handler)
	router.GET("/health", handlers.HealthCheckHandler)
	return router
}
