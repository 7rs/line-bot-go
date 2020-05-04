package bot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dustin/go-humanize"
	sdk "github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/api/youtube/v3"
)

const youtubeDataFlexPath = "resources/flex/youtube.json"

type youtubeData struct {
	Type      string `json:"type"`
	Direction string `json:"direction"`
	Header    struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []struct {
			Type  string `json:"type"`
			Text  string `json:"text"`
			Size  string `json:"size"`
			Align string `json:"align"`
		} `json:"contents"`
	} `json:"header"`
	Hero struct {
		Type        string `json:"type"`
		URL         string `json:"url"`
		Size        string `json:"size"`
		AspectRatio string `json:"aspectRatio"`
		AspectMode  string `json:"aspectMode"`
		Action      struct {
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"action"`
	} `json:"hero"`
	Body struct {
		Type     string `json:"type"`
		Layout   string `json:"layout"`
		Contents []struct {
			Type   string `json:"type"`
			Text   string `json:"text,omitempty"`
			Margin string `json:"margin"`
			Align  string `json:"align,omitempty"`
			Color  string `json:"color,omitempty"`
		} `json:"contents"`
	} `json:"body"`
}

func getYoutubeDataFlexContainer(id string, snippet *youtube.VideoSnippet, statistics *youtube.VideoStatistics) (*sdk.FlexContainer, error) {
	bytes, err := ioutil.ReadFile(youtubeDataFlexPath)
	if err != nil {
		return nil, err
	}

	data := &youtubeData{}
	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	data.Header.Contents[0].Text = snippet.Title
	data.Hero.URL = snippet.Thumbnails.Maxres.Url
	data.Hero.Action.URI = "https://youtu.be/" + id
	data.Body.Contents[0].Text = "▶ " + humanize.Comma(int64(statistics.ViewCount))
	data.Body.Contents[2].Text = "👍 " + humanize.Comma(int64(statistics.LikeCount))

	bytes, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	container, err := sdk.UnmarshalFlexMessageJSON(bytes)
	if err != nil {
		return nil, err
	}

	return &container, nil
}
