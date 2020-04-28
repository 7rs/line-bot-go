package bot

import (
	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
)

type Bot struct {
	Echo *echo.Echo
	API  *sdk.Client
}

func NewBotClient(channelSecret string, channelAccessToken string) (*Bot, error) {
	api, err := sdk.New(channelSecret, channelAccessToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Echo: echo.New(),
		API:  api,
	}, nil
}

func (b *Bot) Start() error {
	SetColog()

	b.SetEcho()
	b.SetEndPoints()

	if err := b.StartEcho(); err != nil {
		return err
	}
	return nil
}
