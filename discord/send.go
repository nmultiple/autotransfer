package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func SendMessage(text, endPoint string) error {
	message := Message {text}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = http.Post(endPoint, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	return nil
}
