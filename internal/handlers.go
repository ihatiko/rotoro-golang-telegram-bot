package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rotoro-golang-telegram-bot/internal/modules"
	telegramClient "rotoro-golang-telegram-bot/pkg/telegram-client"
	"rotoro-golang-telegram-bot/pkg/telegram-client/models"
)

type Handlers struct {
	Client *telegramClient.Client
}

func NewHandlers() *Handlers {
	client := telegramClient.NewClient("")
	return &Handlers{Client: client}
}

// Handler Обработчик запросов
func (s *Handlers) Handler(res http.ResponseWriter, req *http.Request) {
	body := &models.WebHookBody{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	msgBody := modules.Module1Handler(body.Message.Text)
	s.Client.SendResponse(msgBody, body.Message.Chat.ID)
}
