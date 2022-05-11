package internal

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"net/http"
	"rotoro-golang-telegram-bot/internal/modules"
	telegramClient "rotoro-golang-telegram-bot/pkg/telegram-client"
	"rotoro-golang-telegram-bot/pkg/telegram-client/models"
)

type Handlers struct {
	Client *telegramClient.Client
}

func NewHandlers(serverHost string) *Handlers {
	client := telegramClient.NewClient("5378603292:AAHcigJ9ifEZLsycjOALyTz-QHO1cR-Or_g", serverHost, true)
	return &Handlers{Client: client}
}

// Handler Обработчик запросов
func (s *Handlers) Handler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	body := &models.WebHookBody{}
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(body); err != nil {
		log.Error().Err(err).Msg("could not decode request body")
		return
	}
	msgBody := modules.Module1Handler(body.Message.Text)
	s.Client.SendResponse(msgBody, body.Message.Chat.ID)
}
func (s *Handlers) HealthCheckHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		log.Info().Err(err).Msg("get healthCheck")
	}
}
