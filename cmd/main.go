package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var bot *tb.Bot
	var err error

	for {
		bot, err = runBot()
		if err != nil {
			log.Println("Cannot creat and run bot", err)
		} else {
			log.Println("Bot started!")
			break
		}
	}

	bot.Start()
}

// запуск бота
func runBot() (*tb.Bot, error) {
	secret, err := ioutil.ReadFile(".secret/token")
	if err != nil {
		return nil, fmt.Errorf("unable to read token: %v", err)
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:     string(secret),
		Poller:    &tb.LongPoller{Timeout: 2 * time.Second},
		ParseMode: tb.ModeHTML,
	})
	if err != nil {
		return nil, err
	}

	return bot, err
}

// Подключение к базе данных
func RedisDatabase(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       db,
	})

	status := client.Ping(context.TODO())

	if status.Err() != nil {
		return nil, status.Err()
	}

	return client, nil
}
