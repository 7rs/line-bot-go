package bot

import (
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
)

var youtubeLinkRegex = regexp.MustCompile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\?(?:\S*?&?v\=))|youtu\.be\/)([a-zA-Z0-9_-]+)`)

func (b *Bot) getEvents(c echo.Context) ([]*sdk.Event, error) {
	req := c.Request()
	events, err := b.Client.ParseRequest(req)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func handleError(c echo.Context, err error) error {
	if err == sdk.ErrInvalidSignature {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	return c.String(http.StatusInternalServerError, "Internal Server Error")
}

func (b *Bot) sendYoutubeInfo(event *sdk.Event, id string) {
	res, err := b.Service.Videos.List("id,snippet,statistics").Id(id).MaxResults(1).Do()
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	if len(res.Items) == 0 {
		if _, err := b.Client.ReplyMessage(event.ReplyToken, sdk.NewTextMessage("Not found video :/\nVideo ID: "+id)).Do(); err != nil {
			log.Printf("error: %v", err)
		}
		return
	}
	item := res.Items[0]

	container, err := getYoutubeDataFlexContainer(id, item.Snippet, item.Statistics)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	if _, err := b.Client.ReplyMessage(event.ReplyToken, sdk.NewFlexMessage("YouTube", *container)).Do(); err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func (b *Bot) handleMessage(event *sdk.Event) {
	switch msg := event.Message.(type) {
	case *sdk.TextMessage:
		if r := youtubeLinkRegex.FindStringSubmatch(msg.Text); len(r) >= 2 {
			b.sendYoutubeInfo(event, r[1])
		}
	}
}

func (b *Bot) sendWelcomeMessage(event *sdk.Event) {
	for _, member := range event.Joined.Members {
		prof, err := b.Client.GetProfile(member.UserID).Do()
		if err != nil {
			log.Printf("error %v", err)
			continue
		}

		if _, err := b.Client.ReplyMessage(event.ReplyToken, sdk.NewTextMessage("Hi, "+prof.DisplayName)).Do(); err != nil {
			log.Printf("error: %v", err)
			continue
		}
	}
}

func (b *Bot) handleEvents(events []*sdk.Event) {

	for _, event := range events {
		switch event.Type {
		case sdk.EventTypeMessage:
			b.handleMessage(event)
		case sdk.EventTypeMemberJoined:
			b.sendWelcomeMessage(event)
		}
	}
}
