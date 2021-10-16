package telegram

import (
	pocket "github.com/Lapp-coder/go-pocket-sdk"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/config"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	storage      storage.TokenStorage
	pocketClient *pocket.Client
	messages     config.Messages
	redirectURL  string
}

func NewBot(bot *tgbotapi.BotAPI, storage storage.TokenStorage, pocketClient *pocket.Client, messages config.Messages, redirectURL string) *Bot {
	return &Bot{bot: bot, storage: storage, pocketClient: pocketClient, messages: messages, redirectURL: redirectURL}
}

func (b *Bot) Start() error {
	logrus.Infof("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
