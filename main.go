package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/7rs/line-bot-go/bot"
)

func getEnvs() ([]string, error) {
	envs := []string{}

	for _, name := range []string{"CHANNEL_SECRET", "CHANNEL_ACCESS_TOKEN", "API_KEY"} {
		if v := os.Getenv(name); v != "" {
			envs = append(envs, v)
			continue
		}
		return nil, errors.New("Not found " + name)
	}

	return envs, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Not fount .env file.")
	}

	envs, err := getEnvs()
	if err != nil {
		log.Fatal(err)
	}

	b, err := bot.NewBotClient(envs[0], envs[1], envs[2])
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Start(); err != nil {
		log.Fatal(err)
	}
}
