.PHONY: build run docker-build docker-run
build:
	go build -o ./.bin/bot ./cmd/bot/main.go

run:
	./.bin/bot

docker-build:
	docker build -t pocketer-telegram-bot .

docker-run:
	docker run --name pocketer-telegram-bot-container --env-file .env pocketer-telegram-bot

.DEFAULT_GHOAL: build
