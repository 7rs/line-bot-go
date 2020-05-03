package bot

import (
	"log"
	"net/http"
	"regexp"
	"strings"

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

func (b *Bot) sendYoutubeInfo(event *sdk.Event, text string) {
	queries := youtubeLinkRegex.FindStringSubmatch(text)
	if len(queries) < 2 {
		log.Printf("warn: Not found id.")
		return
	}

	res, err := b.Service.Videos.List("id,snippet,statistics").Id(queries[1]).MaxResults(1).Do()
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	if len(res.Items) == 0 {
		log.Printf("warn: Not found video.")
		return
	}
	item := res.Items[0]

	info := strings.Join([]string{
		"Title: " + item.Snippet.Title,
		"Views: " + string(item.Statistics.ViewCount),
		"Likes: " + string(item.Statistics.LikeCount),
	}, "\n")

	if _, err := b.Client.ReplyMessage(event.ReplyToken, sdk.NewTextMessage(info)).Do(); err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func (b *Bot) handleMessage(event *sdk.Event) {
	switch msg := event.Message.(type) {
	case *sdk.TextMessage:
		text := strings.ToLower(msg.Text)
		if strings.Contains(text, "youtube.com") {
			b.sendYoutubeInfo(event, text)
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
