package telegram_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rotoro-golang-telegram-bot/pkg/telegram-client/models"
)

type Client struct {
	Path string
}

func NewClient(apiKey string) *Client {
	url := fmt.Sprintf("%s%s%s", telegramUrl, apiKey, webhookPath)
	return &Client{Path: url}
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
		c.Path,
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
