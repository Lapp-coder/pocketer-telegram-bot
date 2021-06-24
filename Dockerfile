FROM golang:1.16.4-alpine3.13 AS builder

COPY . /github.com/Lapp-coder/pocketer-telegram-bot/
WORKDIR /github.com/Lapp-coder/pocketer-telegram-bot/

RUN go mod download
RUN go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Lapp-coder/pocketer-telegram-bot/.bin/bot .
COPY --from=0 /github.com/Lapp-coder/pocketer-telegram-bot/configs configs/

CMD ["./bot"]
