package main

import (
	"log"

	"github.com/dreddsa5dies/gobotredis/bot"
	"github.com/dreddsa5dies/gobotredis/storage"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var b *tb.Bot
	var err error

	go storage.UpdatePair()

	for {
		b, err = bot.Run()
		if err != nil {
			log.Println("[!] Cannot creat and run bot", err)
		} else {
			log.Println("Bot started!")
			break
		}
	}

	b.Start()
}
