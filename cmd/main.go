package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc/grpclog"
	"tinkoff-investment-bot/internal/bot/listener"
	"tinkoff-investment-bot/internal/bot/model"
	"tinkoff-investment-bot/internal/connect"
	ms "tinkoff-investment-bot/internal/model/settings"
)

func main() {
	botAPI, err := tgbotapi.NewBotAPI("5844971569:AAG08-g6E4ypD_XvIb2wEjYRjDo0owdMyhw")
	if err != nil {
		grpclog.Fatalln("connect Telegram bot: failed to create bot: %v", err)
	}

	tinkoffInvestBot := model.New(botAPI)

	settings := ms.NewSettings()

	defer func() {
		connect.Close(&settings)
	}()

	listener.ListenUpdates(tinkoffInvestBot, &settings)
}
