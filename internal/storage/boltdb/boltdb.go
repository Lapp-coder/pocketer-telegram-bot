package boltdb

import (
	"strconv"

	"github.com/Lapp-coder/pocketer-telegram-bot/internal/storage"
	"github.com/boltdb/bolt"
)

type TokenStorage struct {
	db *bolt.DB
}

func NewTokenStorage(db *bolt.DB) *TokenStorage {
	return &TokenStorage{db: db}
}

func (s *TokenStorage) Save(chatID int64, token string, bucket storage.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}
func (s *TokenStorage) Get(chatID int64, bucket storage.Bucket) (string, error) {
	var token []byte

	if err := s.db.View(func(tx *bolt.Tx) error {
		txBucket := tx.Bucket([]byte(bucket))
		token = txBucket.Get(intToBytes(chatID))
		return nil
	}); err != nil {
		return "", err
	}

	if len(token) == 0 {
		return "", errTokenNotFound
	}

	return string(token), nil
}

func (s *TokenStorage) Delete(chatID int64, bucket storage.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		txBucket := tx.Bucket([]byte(bucket))
		return txBucket.Delete(intToBytes(chatID))
	})
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
