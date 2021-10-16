FROM golang:1.17-alpine3.14 AS builder

COPY . /github.com/Lapp-coder/pocketer-telegram-bot/
WORKDIR /github.com/Lapp-coder/pocketer-telegram-bot/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -installsuffix "static" -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Lapp-coder/pocketer-telegram-bot/.bin/bot .
COPY --from=builder /github.com/Lapp-coder/pocketer-telegram-bot/configs configs/

CMD ["./bot"]
