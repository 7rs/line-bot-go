package bot

import (
	"log"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	sdk "github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/api/youtube/v3"
)

var youtubeLinkRegex = regexp.MustCompile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\?(?:\S*?&?v\=))|youtu\.be\/)([a-zA-Z0-9_-]+)`)

func (p *Client) getEvents(c echo.Context) ([]*sdk.Event, error) {
	req := c.Request()
	events, err := p.Client.ParseRequest(req)
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

func (p *Client) searchVideoWithID(event *sdk.Event, id string) (*youtube.Video, error) {
	res, err := p.Service.Videos.List("id,snippet,statistics").Id(id).MaxResults(1).Do()
	if err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		if _, err := p.Client.ReplyMessage(event.ReplyToken, sdk.NewTextMessage("Not found video :/\nVideo ID: "+id)).Do(); err != nil {
			return nil, err
		}
	}

	return res.Items[0], nil
}

func (p *Client) sendYoutubeInfo(event *sdk.Event, id string) error {
	video, err := p.searchVideoWithID(event, id)
	if err != nil {
		return err
	}

	container, err := getYoutubeDataFlexContainer(id, video.Snippet, video.Statistics)
	if err != nil {
		return err
	}

	if _, err := p.Client.ReplyMessage(event.ReplyToken, sdk.NewFlexMessage("YouTube", *container)).Do(); err != nil {
		return err
	}

	return nil
}

func (p *Client) sendWelcomeMessage(event *sdk.Event) {
	for _, member := range event.Joined.Members {
		prof, err := p.Client.GetProfile(member.UserID).Do()
		if err != nil {
			log.Printf("error %v", err)
			continue
		}

		if _, err := p.Client.ReplyMessage(event.ReplyToken, sdk.NewTextMessage("Hi, "+prof.DisplayName)).Do(); err != nil {
			log.Printf("error: %v", err)
			continue
		}
	}
}

func (p *Client) handleMessage(event *sdk.Event) {
	switch msg := event.Message.(type) {
	case *sdk.TextMessage:
		if r := youtubeLinkRegex.FindStringSubmatch(msg.Text); len(r) >= 2 {
			if err := p.sendYoutubeInfo(event, r[1]); err != nil {
				log.Printf("error: %v", err)
			}
		}
	}
}

func (p *Client) handleEvents(events []*sdk.Event) {
	for _, event := range events {
		switch event.Type {
		case sdk.EventTypeMessage:
			p.handleMessage(event)
		case sdk.EventTypeMemberJoined:
			p.sendWelcomeMessage(event)
		}
	}
}
