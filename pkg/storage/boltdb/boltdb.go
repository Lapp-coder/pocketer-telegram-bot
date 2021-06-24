package boltdb

import (
	"errors"
	"github.com/Lapp-coder/pocketer-telegram-bot/pkg/storage"
	"github.com/boltdb/bolt"
	"strconv"
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
	var token string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}

func (s *TokenStorage) Delete(chatID int64, bucket storage.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete(intToBytes(chatID))
	})
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
