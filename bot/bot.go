package bot

import (
	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Bot struct {
	Echo    *echo.Echo
	Client  *sdk.Client
	Service *youtube.Service
}

func NewBotClient(channelSecret string, channelAccessToken string, apiKey string) (*Bot, error) {
	sdkClient, err := sdk.New(channelSecret, channelAccessToken)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &Bot{
		Echo:    echo.New(),
		Client:  sdkClient,
		Service: service,
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
