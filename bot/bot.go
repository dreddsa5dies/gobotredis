package bot

import (
	"fmt"
	"io/ioutil"
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

	bot.Handle(tb.OnText, func(m *tb.Message) {
		bot.Send(m.Sender, "hello world")
	})

	return bot, err
}
