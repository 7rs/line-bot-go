package bot

import "github.com/7rs/line-bot-go/line"

type BotClient struct {
	Api *line.APIClient
}

func NewBotClient(api *line.APIClient) *BotClient {
	return &BotClient{Api: api}
}
