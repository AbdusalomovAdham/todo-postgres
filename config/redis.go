package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("Error connected RedisDb")
	}
	log.Println("Redis connected seccessfuly !")
	return client
}
