package connect

import (
	"context"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinkoff-investment-bot/internal/connect/config"
	"tinkoff-investment-bot/internal/model"
	printbot "tinkoff-investment-bot/internal/print-bot"
	"tinkoff-investment-bot/internal/storage"
)

func Config(telegramID string) (*investgo.Client, *gorm.DB, context.CancelFunc, *zap.SugaredLogger, *model.Tracker) {
	logger := getLogger()

	db, err := connectDB()
	if err != nil {
		logger.Errorf(err.Error())
	}

	userStorage := storage.NewUserStorage(db)

	user := userStorage.GetUserByTelegramID(telegramID)

	var token string

	if user.ID != 0 {
		token = user.Token
	} else {
		token, err = printbot.GetTokenFromUser()
		if err != nil {
			logger.Errorf(err.Error())
		}
	}

	client, cancel, err := config.ClientTinkoffInvestByConfigYaml(logger, &token)
	if err != nil {
		logger.Errorf(err.Error())
	}

	err = userStorage.AddUser(&model.User{TelegramID: telegramID, Token: token})
	if err != nil {
		logger.Errorf(err.Error())
	}

	var tracker model.Tracker
	tracker.NewTracker(client)

	return client, db, cancel, logger, &tracker
}
