package storage

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/dreddsa5dies/gobotredis/getpair"
	"github.com/go-redis/redis/v8"
)

// Структура значений валютных пар
type CUR struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		Rub float64 `json:"RUB"`
		Usd float64 `json:"USD"`
	} `json:"rates"`
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

	log.Println("[*] redis connect successfully")

	return client, nil
}

// Запись данных в базу
func SetData(data []byte, key string) (err error) {
	client, err := RedisDatabase(0)
	if err != nil {
		return err
	}

	d, err := getpair.GetCur()
	if err != nil {
		return err
	}

	err = client.Set(context.TODO(), key, d, 0).Err()
	if err != nil {
		return err
	}

	log.Println("[*] set data to redis successfully")

	return nil
}

// Получение данных по ключу (текущий день)
func GetData(key string) (data *CUR, err error) {
	client, err := RedisDatabase(0)
	if err != nil {
		return nil, err
	}

	val, err := client.Get(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}

	var d *CUR

	json.Unmarshal([]byte(val), d)

	log.Println("[*] get data form redis successfully")

	return d, nil
}

// Сохранение значений валютных пар и их обновление 2 раза в день
func UpdatePair() {
	for {
		currentTime := time.Now()
		key := currentTime.Format("09-07-2017")

		pair, err := getpair.GetCur()
		if err != nil {
			log.Println(err)
		}
		err = SetData(pair, key)
		if err != nil {
			log.Println(err)
		}

		log.Println("[*] pair update successfully")

		time.Sleep(12 * time.Hour)
	}
}
