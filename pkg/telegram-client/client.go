package telegram_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"rotoro-golang-telegram-bot/pkg/telegram-client/models"
)

type Client struct {
	Path                 string
	TelegramRegistryHost string
	AutoRegistry         bool
}

func NewClient(apiKey, telegramRegistryHost string, autoRegistry bool) *Client {
	url := fmt.Sprintf("%s%s%s", telegramUrl, apiKey)
	client := &Client{Path: url, AutoRegistry: autoRegistry, TelegramRegistryHost: telegramRegistryHost}
	if client.AutoRegistry {
		client.registryWebHook()
	}
	return client
}

func (c *Client) registryWebHook() {
	request := models.RegistryWebHookRequest{Url: c.TelegramRegistryHost}
	body, err := json.Marshal(&request)
	if err != nil {
		log.Error().Err(err).Msg("registryWebHook")
		return
	}
	res, err := http.Post(
		fmt.Sprintf("%s%s", c.Path, setWebHook),
		applicationJsonHeader,
		bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msg("registryWebHook http.Post")
	}
	if res.StatusCode != http.StatusOK {
		log.Error().Err(err).Msgf("unexpected status %s", res.Status)
	}

	var buffer []byte
	_, err = res.Body.Read(buffer)
	if err != nil {
		log.Error().Err(err).Msg("res.Body.Read(buffer)")
		return
	}
	response := &models.RegistryWebHookResponse{}
	err = json.Unmarshal(buffer, response)
	if err != nil {
		log.Error().Err(err).Msg("json.Unmarshal(buffer, response)")
		return
	}
	log.Info().Msgf("status %b description %s result %b", response.Ok, response.Description, response.Result)
}
func (c *Client) SendResponse(message string, chatID int64) error {
	reqBody := &models.SendMessageBody{
		ChatID: chatID,
		Text:   message,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	res, err := http.Post(
		fmt.Sprintf("%s%s", c.Path, sendMessage),
		applicationJsonHeader,
		bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("unexpected status %s", res.Status))
	}
	return nil
}
