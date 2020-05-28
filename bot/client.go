package bot

import (
	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var err error
var e = echo.New()

type Client struct {
	Client  *sdk.Client
	Service *youtube.Service
}

func NewBotClient(channelSecret string, channelAccessToken string, apiKey string) (*Client, error) {
	sdkClient, err := sdk.New(channelSecret, channelAccessToken)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: sdkClient,
	}, nil
}

func (p *Client) LoginYoutubeService(apiKey string) error {
	ctx := context.Background()
	p.Service, err = youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}

	return nil
}

func (p *Client) Start() error {
	SetColog()

	SetEcho()
	p.SetEndPoints()

	if err := StartEcho(); err != nil {
		return err
	}
	return nil
}
