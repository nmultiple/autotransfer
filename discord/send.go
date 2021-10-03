package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	LimitTooShort = errors.New("Limit length is too short to split")
)

const messageLengthLimit = 2000

type Message struct {
	Content string `json:"content"`
}

func splitToLimit(text string, limit int) ([]string, error) {
	if limit <= 50 {
		return nil, LimitTooShort
	}

	sourceText := []rune(text)
	var result []string

	for {
		var right int
		if limit-50 >= len(sourceText) {
			right = len(sourceText)
		} else {
			right = limit - 50
		}

		result = append(result, string(sourceText[:right]))
		sourceText = sourceText[right:]

		if len(sourceText) == 0 {
			break
		}
	}

	return result, nil
}

func postMessage(text, endPoint string) error {
	message := Message{text}

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

func SendMessage(text, endPoint string) error {
	partialMessages, err := splitToLimit(text, messageLengthLimit)
	if err != nil {
		return err
	}

	if len(partialMessages) == 1 {
		return postMessage(partialMessages[0], endPoint)
	} else {
		var err error
		for i, msg := range partialMessages {
			newText := fmt.Sprintf("%v\n---続きます (%v/%v)", msg, i+1, len(partialMessages))
			if err1 := postMessage(newText, endPoint); err1 != nil {
				err = err1
			}
		}

		return err
	}
}
