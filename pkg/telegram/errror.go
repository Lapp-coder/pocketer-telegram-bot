package telegram

import "errors"

var (
	errInvalidURL               = errors.New("url is invalid")
	errUnauthorized             = errors.New("user is not authorized")
	errFailedToSave             = errors.New("failed to save")
	errFailedToGet              = errors.New("failed to get")
	errFailedToAuthorized       = errors.New("failed to authorize user")
	errFailedToGenerateAuthLink = errors.New("failed to generate authorization link")
)
