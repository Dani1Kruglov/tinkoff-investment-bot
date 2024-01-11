package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tinkoff-investment-bot/internal/bot/listener"
	"tinkoff-investment-bot/internal/bot/model"
	"tinkoff-investment-bot/internal/connect"
	ms "tinkoff-investment-bot/internal/model/settings"
)

func main() {
	botAPI, err := tgbotapi.NewBotAPI("TELEGRAM_BOT_API_TOKEN")
	if err != nil {
		fmt.Errorf("failed to create bot: %v", err)
	}

	tinkoffInvestBot := model.New(botAPI)

	settings := ms.NewSettings()

	defer func() {
		connect.Close(&settings)
	}()

	listener.ListenUpdates(tinkoffInvestBot, &settings)
}
