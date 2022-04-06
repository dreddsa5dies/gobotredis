package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

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
