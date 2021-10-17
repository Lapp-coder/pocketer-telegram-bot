.PHONY: build run docker-build docker-run
build:
	go build -o ./.bin/bot ./cmd/bot/main.go

run: build
	./.bin/bot

docker-build:
	docker build -t pocketer-telegram-bot .

docker-run: docker-build
	docker run --name pocketer-telegram-bot --env-file .env pocketer-telegram-bot

.DEFAULT_GHOAL: build
