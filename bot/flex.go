package bot

import (
	"io/ioutil"
	"strconv"

	"github.com/bitly/go-simplejson"
	sdk "github.com/line/line-bot-sdk-go/linebot"
)

const youtubeDataFlexPath = "resources/flex/youtube.json"

type youtubeData struct {
	Thumbnail string
	Title     string
	Views     int
	Likes     int
	ID        string
}

func loadingFlex(name string) (*simplejson.Json, error) {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	flex, err := simplejson.NewJson(bytes)
	if err != nil {
		return nil, err
	}

	return flex, nil
}

func setYoutubeData(flex *simplejson.Json, data *youtubeData) (*simplejson.Json, error) {
	contents, err := flex.GetPath("body", "contents").Array()
	if err != nil {
		return nil, err
	}

	thumbnailComponent := contents[0].(map[string]interface{})
	thumbnailComponent["url"] = data.Thumbnail

	titleBoxComponent := contents[1].(map[string]interface{})
	titleComponent := titleBoxComponent["contents"].([]interface{})[0].(map[string]interface{})
	titleComponent["text"] = data.Title

	insightBoxComponents := contents[2].(map[string]interface{})["contents"].([]interface{})

	viewsBoxComponent := insightBoxComponents[0].(map[string]interface{})
	viewsComponent := viewsBoxComponent["contents"].([]interface{})[0].(map[string]interface{})
	viewsComponent["text"] = "▶ " + strconv.Itoa(data.Views)

	likesBoxComponent := insightBoxComponents[1].(map[string]interface{})
	likesComponent := likesBoxComponent["contents"].([]interface{})[0].(map[string]interface{})
	likesComponent["text"] = "👍 " + strconv.Itoa(data.Likes)

	action, err := flex.GetPath("body", "action").Map()
	if err != nil {
		return nil, err
	}
	action["uri"] = "https://youtu.be/" + data.ID

	flex.SetPath([]string{"body", "contents"}, contents)
	flex.SetPath([]string{"body", "action"}, action)

	return flex, nil
}

func getYoutubeDataFlexContainer(data *youtubeData) (*sdk.FlexContainer, error) {
	flex, err := loadingFlex(youtubeDataFlexPath)
	if err != nil {
		return nil, err
	}

	flex, err = setYoutubeData(flex, data)
	if err != nil {
		return nil, err
	}

	r, err := flex.Encode()
	if err != nil {
		return nil, err
	}

	container, err := sdk.UnmarshalFlexMessageJSON(r)
	if err != nil {
		return nil, err
	}

	return &container, nil
}
