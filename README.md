[![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/gobotredis)](https://goreportcard.com/report/github.com/dreddsa5dies/gobotredis) ![License](https://img.shields.io/badge/License-MIT-blue.svg)  

## Создание бота telegram с хранением в redis  
- [x] развертывание redis в docker
- [x] развертывания min go bot telegram
- [x] profit


## [Bot](http://t.me/testgoredis_bot)

### Description
Тестовый бот для выдачи информации о курсах валют

Бот регистрируется на [@BotFather](https://t.me/botfather)

## Packages
Use [Go Modules](https://blog.golang.org/using-go-modules)

```bash
docker pull redis
```

## Start

### Токены
Сохраняются в папку ./secret/ проекта в виде файлов формата:
- ./secret/token

Токен telegram для активации бота
- ./secret/fixer

Ключ API [fixer](https://fixer.io/usage)

Все токены записаны обычной одной строкой.

### ВАЖНО!

Указать папку .secret/ в .gitignore

### Старт проекта

```bash
    docker run --name redis-test-instance -p 6379:6379 -d redis
    docker ps
    cd ../cmd
    go run main.go
  ```
## The code contains comments in Russian

## License
This project is licensed under MIT license. Please read the [LICENSE](https://github.com/dreddsa5dies/gobotredis/tree/master/LICENSE.md) file.  