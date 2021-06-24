.PHONY: build run build-container run-container
build:
	go build -o ./.bin/bot ./cmd/bot/main.go

run:
	./.bin/bot

build-container:
	docker build -t pocketer-telegram-bot .

run-container:
	docker run --name pocketer-telegram-bot-container --env-file .env pocketer-telegram-bot

.DEFAULT_GHOAL: build
