package Telegram

import (
	"context"
	"log"

	"github.com/E7ast1c/Mao/ChatGPT"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBot struct {
	gptClient        ChatGPT.Client
	permittedUserVal string
	bot              *tgbotapi.BotAPI
}

func New(tgToken string, permittedUserVal string, gptClient ChatGPT.Client, debug bool,
) (*TgBot, error) {

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = debug
	log.Printf("Bot autorized as %s", bot.Self.UserName)
	return &TgBot{bot: bot, permittedUserVal: permittedUserVal, gptClient: gptClient}, nil
}

func (t *TgBot) send(chatId int64, content string) {
	_, err := t.bot.Send(tgbotapi.NewMessage(chatId, content))
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}
}

func (t *TgBot) ReceiveMessages() error {
	updates, _ := t.bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Timeout: 60,
	})

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "":
			t.send(update.Message.Chat.ID, "empty request")
		case "/start":
			t.send(update.Message.Chat.ID, "let`s get started")
		default:
			if t.permittedUserVal != update.Message.Chat.UserName {
				t.send(update.Message.Chat.ID, "access restricted")
				break
			}

			t.send(update.Message.Chat.ID, "\U000023F1")
			resp, err := t.gptClient.ChatCompletionRequest(context.Background(), update.Message.Text)
			if err != nil {
				t.send(update.Message.Chat.ID, err.Error())
				break
			}
			t.send(update.Message.Chat.ID, resp)
		}
	}
	return nil
}
