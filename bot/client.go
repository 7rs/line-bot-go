package bot

import (
	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var e = echo.New()
var ctx = context.Background()

type Tokens struct {
	ChannelSecret      string
	ChannelAccessToken string
	ApiKey             string
}

type Client struct {
	Client  *sdk.Client
	Service *youtube.Service
}

func NewBotClient(tokens *Tokens) (*Client, error) {
	sdkClient, err := sdk.New(tokens.ChannelSecret, tokens.ChannelAccessToken)
	if err != nil {
		return nil, err
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(tokens.ApiKey))
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  sdkClient,
		Service: service,
	}, nil
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
