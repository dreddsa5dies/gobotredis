package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	log.Println("Starting... ok!")

	var bot *tb.Bot
	var err error

	for {
		bot, err = runBot()
		if err != nil {
			log.Println("Cannot creat and run bot", err)
		} else {
			log.Println("Bot started")
			break
		}
	}
	bot.Start()
}

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
