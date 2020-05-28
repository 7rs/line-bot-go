package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/7rs/line-bot-go/bot"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warn: Not fount .env file.")
	}

	b, err := bot.NewBotClient(&bot.Tokens{
		ChannelSecret:      os.Getenv("CHANNEL_SECRET"),
		ChannelAccessToken: os.Getenv("CHANNEL_ACCESS_TOKEN"),
		ApiKey:             os.Getenv("API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Start(); err != nil {
		log.Fatal(err)
	}
}
