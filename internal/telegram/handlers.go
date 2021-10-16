package telegram

import (
	"context"
	pocket "github.com/Lapp-coder/go-pocket-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/url"
)

const (
	commandStart = "start"
	commandGet   = "get"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandGet:
		return b.handleGetCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.SavedSuccessfully)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		requestToken, err := b.getRequestToken(message.Chat.ID)
		if err != nil {
			return errUnauthorized
		}

		accessToken, err = b.userAuthentication(message.Chat.ID, requestToken)
		if err != nil {
			return errFailedToAuthorized
		}
	}

	if err = b.pocketClient.Add(context.Background(), pocket.AddInput{AccessToken: accessToken, URL: message.Text}); err != nil {
		return errFailedToSave
	}

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message.Chat.ID)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.AlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleGetCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		requestToken, err := b.getRequestToken(message.Chat.ID)
		if err != nil {
			return errUnauthorized
		}

		accessToken, err = b.userAuthentication(message.Chat.ID, requestToken)
		if err != nil {
			return errFailedToAuthorized
		}
	}

	items, err := b.pocketClient.Retrieving(context.Background(), pocket.RetrievingInput{AccessToken: accessToken})
	if err != nil {
		return errFailedToGet
	}

	for _, item := range items {
		msg.Text = item.GivenUrl
		if _, err = b.bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Errors.Default)

	switch err {
	case errInvalidURL:
		msg.Text = b.messages.Errors.InvalidURL
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = b.messages.Errors.Unauthorized
		b.bot.Send(msg)
	case errFailedToSave:
		msg.Text = b.messages.Errors.FailedToSave
		b.bot.Send(msg)
	case errFailedToAuthorized:
		msg.Text = b.messages.Errors.FailedToAuthorized
		b.bot.Send(msg)
	case errFailedToGenerateAuthLink:
		msg.Text = b.messages.Errors.FailedToGenerate
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
