package utils

import (
	"context"
	"myproject/config"
	"time"
)

var rdb = config.ConnectRDB()

func Set(key string, value interface{}, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rdb.Set(ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rdb.Get(ctx, key).Result()
}

func Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rdb.Del(ctx, key).Err()
}
