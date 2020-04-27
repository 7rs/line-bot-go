package line

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type replyMessageBody struct {
	ReplyToken           string      `json:"replyToken"`
	Messages             interface{} `json:"messages"`
	NotificationDisabled bool        `json:"notificationDisabled"`
}

func NewTextMessage(text string) map[string]interface{} {
	return map[string]interface{}{
		"type": "text",
		"text": text,
	}
}

func (p *APIClient) SendReplyMessage(replyToken string, messages []map[string]interface{}, notificationDisabled bool) error {
	body, err := json.Marshal(&replyMessageBody{
		ReplyToken:           replyToken,
		Messages:             messages,
		NotificationDisabled: notificationDisabled,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", APIHost+Reply, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.Token)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
