package main

import (
	"log"

	pocket "github.com/Lapp-coder/go-pocket-sdk"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/config"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/storage"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/storage/boltdb"
	"github.com/Lapp-coder/pocketer-telegram-bot/internal/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	fsModeReadWriteOnly = 0600
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	tokenStorage := boltdb.NewTokenStorage(db)

	telegramBot := telegram.NewBot(bot, tokenStorage, pocketClient, cfg.Messages, cfg.RedirectURL)
	if err = telegramBot.Start(); err != nil {
		log.Fatalf("an error occurred when starting telegram bot: %s", err)
	}
}

func initDB(dbPath string) (*bolt.DB, error) {
	db, err := bolt.Open(dbPath, fsModeReadWriteOnly, nil)
	if err != nil {
		return nil, err
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists([]byte(storage.AccessTokens)); err != nil {
			return err
		}

		if _, err = tx.CreateBucketIfNotExists([]byte(storage.RequestTokens)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
