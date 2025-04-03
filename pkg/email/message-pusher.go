package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/9688101/HX/config"
)

type request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Channel     string `json:"channel"`
	Token       string `json:"token"`
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SendMessage(title, description, content string) error {
	cfg := config.GetMailConfig()
	if cfg.SMTPFrom == "" {
		return errors.New("message pusher address is not set")
	}

	req := request{
		Title:       title,
		Description: description,
		Content:     content,
		Token:       cfg.SMTPToken,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(cfg.SMTPFrom, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	if !res.Success {
		return errors.New(res.Message)
	}
	return nil
}
