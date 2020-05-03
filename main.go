package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/7rs/line-bot-go/bot"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Not fount .env file.")
	}

	vars := []string{}

	for _, name := range []string{"CHANNEL_SECRET", "CHANNEL_ACCESS_TOKEN", "API_KEY"} {
		if v := os.Getenv(name); v != "" {
			vars = append(vars, v)
			continue
		}
		log.Fatalf("Not found %v", name)
	}

	b, err := bot.NewBotClient(vars[0], vars[1], vars[2])
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Start(); err != nil {
		log.Fatal(err)
	}
}
