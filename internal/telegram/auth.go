package telegram

import (
	"context"
	"fmt"

	"github.com/Lapp-coder/pocketer-telegram-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) getAccessTokenIfAuthorized(chatID int64) (string, error) {
	accessToken, err := b.getAccessToken(chatID)
	if err != nil {
		requestToken, err := b.getRequestToken(chatID)
		if err != nil {
			return "", errUnauthorized
		}

		accessToken, err = b.userAuthentication(chatID, requestToken)
		if err != nil {
			return "", errFailedToAuthorized
		}
	}

	return accessToken, nil
}

func (b *Bot) userAuthentication(chatID int64, requestToken string) (string, error) {
	auth, err := b.pocketClient.Authorize(context.Background(), requestToken)
	if err != nil {
		return "", err
	}

	if err = b.storage.Save(chatID, auth.AccessToken, storage.AccessTokens); err != nil {
		return "", err
	}

	return auth.AccessToken, nil
}

func (b *Bot) initAuthorizationProcess(chatID int64) error {
	authLink, err := b.generateAuthorizationLink(chatID)
	if err != nil {
		return errFailedToGenerateAuthLink
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(b.messages.Responses.Start, authLink))
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.storage.Get(chatID, storage.AccessTokens)
}

func (b *Bot) getRequestToken(chatID int64) (string, error) {
	return b.storage.Get(chatID, storage.RequestTokens)
}

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), b.redirectURL, "")
	if err != nil {
		return "", err
	}

	if err = b.storage.Save(chatID, requestToken, storage.RequestTokens); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken)
}
