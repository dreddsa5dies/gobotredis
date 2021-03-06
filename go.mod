module github.com/dreddsa5dies/gobotredis

go 1.17

require (
	github.com/go-redis/redis/v8 v8.11.5
	gopkg.in/tucnak/telebot.v2 v2.5.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/pkg/errors v0.8.1 // indirect
)

replace bot => ../bot

replace getrair => ../getrair

replace storage => ../storage
