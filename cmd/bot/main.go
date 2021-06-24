package main

import (
	pocket "github.com/Lapp-coder/go-pocket-sdk"
	"github.com/Lapp-coder/pocketer-telegram-bot/pkg/config"
	"github.com/Lapp-coder/pocketer-telegram-bot/pkg/storage"
	"github.com/Lapp-coder/pocketer-telegram-bot/pkg/storage/boltdb"
	"github.com/Lapp-coder/pocketer-telegram-bot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	stor := boltdb.NewTokenStorage(db)

	telegramBot := telegram.NewBot(bot, stor, pocketClient, cfg.Messages, cfg.RedirectURL)
	if err = telegramBot.Start(); err != nil {
		log.Fatalf("an error occurred when starting telegram bot: %s", err)
	}
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists([]byte(storage.AccessTokens)); err != nil {
			return err
		}

		if _, err = tx.CreateBucketIfNotExists([]byte(storage.RequestTokens)); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
