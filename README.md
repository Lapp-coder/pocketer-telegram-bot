# Pocketer bot [![Go](https://img.shields.io/badge/go-1.17-blue)](https://golang.org/doc/go1.17) [![Release](https://img.shields.io/badge/release-1.1.0-success)](https://github.com/Lapp-coder/pocketer-telegram-bot/releases)
![image](images/pocketer-telegram-bot.jpeg)
## [Pocketer](https://t.me/PocketerBot) - a client for the [Pocket](https://getpocket.com) service in Telegram

***

### Features
* Quickly saving a link to Pocket
* Getting all saved links from Pocket
* Deleting link by id

### Installation and running
***
To install Pocketer, run to next command:
```
$ git clone github.com/Lapp-coder/pocketer-telegram-bot
```

After running command:
```
$ cd pocketer-telegram-bot && touch .env
```

Next, specify the contents of the .env file in this format:
```dotenv
TELEGRAM_BOT_TOKEN=<your-token>
POCKET_CONSUMER_KEY=<your-key>
```

Finally, to run the bot in the docker, run the following command:
```
$ docker build -t pocketer-telegram-bot . && \
  docker run --name pocketer-telegram-bot --env-file .env pocketer-telegram-bot
```

Or, if you have an installed utility "make":
```
$ make docker-run
```
