package main

import (
	"context"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var link = "https://www.youtube.com/watch?v=-f0BhznUNAw&list=RD-f0BhznUNAw&start_radio=1"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Not fount .env file.")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	var regex = regexp.MustCompile(`(?:youtube\.com\/\S*(?:(?:\/e(?:mbed))?\/|watch\?(?:\S*?&?v\=))|youtu\.be\/)([a-zA-Z0-9_-]+)`)
	queries := regex.FindStringSubmatch(link)
	if len(queries) < 2 {
		log.Fatal("Invaild youtube link.")
	}

	log.Printf("Video ID: %v", queries[1])
	res, err := service.Videos.List("id,snippet,statistics").Id(queries[1]).MaxResults(1).Do()
	if err != nil {
		log.Fatal(err)
	}

	if len(res.Items) == 0 {
		log.Printf("warn: Not found video.")
		return
	}
	item := res.Items[0]

	log.Println("Video Title: " + item.Snippet.Title)

	log.Println("Video Views: " + strconv.Itoa(int(item.Statistics.ViewCount)))
	log.Println("Video Likes: " + strconv.Itoa(int(item.Statistics.LikeCount)))
}
