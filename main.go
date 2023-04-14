package main

import (
	"log"
	"os"

	"github.com/E7ast1c/Mao/ChatGPT"
	"github.com/E7ast1c/Mao/Telegram"
)

const (
	tgBotToken    = "TG_BOT_TOKEN"
	gptToken      = "GPT_TOKEN"
	permittedUser = "PERMITTED_USER"
)

func main() {
	permittedUserVal := os.Getenv(permittedUser)
	if permittedUserVal == "" {
		log.Fatalf("%s not found in environment variable", permittedUser)
	}

	tgBotTokenVal := os.Getenv(tgBotToken)
	if tgBotTokenVal == "" {
		log.Fatalf("%s not found in environment variable", tgBotToken)
	}

	gptTokenVal := os.Getenv(gptToken)
	if gptTokenVal == "" {
		log.Fatalf("%s not found in environment variable", gptToken)
	}

	gptClient := ChatGPT.New(gptTokenVal)
	tgBot, err := Telegram.New(tgBotTokenVal, permittedUserVal, gptClient, true)
	if err != nil {
		log.Fatal(err)
	}

	err = tgBot.ReceiveMessages()
	if err != nil {
		log.Fatalf("ReceiveMessages, error:%s", err)
	}
}
