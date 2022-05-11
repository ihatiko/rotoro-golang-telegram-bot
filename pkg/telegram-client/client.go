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
	url := fmt.Sprintf("%s%s", telegramUrl, apiKey)
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
	path := fmt.Sprintf("%s%s", c.Path, setWebHook)
	res, err := http.Post(
		path,
		applicationJsonHeader,
		bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msg("registryWebHook http.Post")
	}
	if res.StatusCode != http.StatusOK {
		log.Error().Err(err).Msgf("unexpected status %s", res.Status)
	}

	response := &models.RegistryWebHookResponse{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
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
