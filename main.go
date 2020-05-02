package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/7rs/line-bot-go/bot"
)

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadDotEnv()
	b, err := bot.NewBotClient(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Start(); err != nil {
		log.Fatal(err)
	}
}
