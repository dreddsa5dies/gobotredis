package bot

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Запуск бота
func Run() (*tb.Bot, error) {
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

	bot.SetCommands([]tb.Command{
		{Text: "/start", Description: "start the bot"},
	})

	bot.Handle("/start", func(m *tb.Message) {
		bot.UnpinAll(m.Chat)
		user := m.Sender
		bot.Delete(m)

		selectorLocale := &tb.ReplyMarkup{}
		RuBtn := selectorLocale.Data("€ EUR", "eur_btn", "ru")
		EnBtn := selectorLocale.Data("$ USD", "usd_btn", "en")
		selectorLocale.Inline(selectorLocale.Row(RuBtn, EnBtn))
		bot.Send(user, "Выбор валюты", selectorLocale)
		bot.Handle(&RuBtn, func(c *tb.Callback) {
			user := m.Sender
			log.Printf("DEBUG: user %v request RUB/EUR", user.Username)
			bot.Delete(c.Message)
			bot.Send(user, "Цену пока не знаю")
			bot.Respond(c)
		})
		bot.Handle(&EnBtn, func(c *tb.Callback) {
			user := m.Sender
			log.Printf("DEBUG: user %v request RUB/USD", user.Username)
			bot.Delete(c.Message)
			bot.Send(user, "Цену пока не знаю")
			bot.Respond(c)
		})
	})

	return bot, err
}
