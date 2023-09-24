package redis

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Setup() {
	getInstance()
}

func getInstance() *redis.Client {
	if client == nil {
		return connect()
	}
	return client
}

func connect() *redis.Client {
	dbStr := os.Getenv("REDIS_DATABASE")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		if dbStr == "" {
			db = 0
		} else {
			panic("REDIS_DATABASE env var is not a number")
		}
	}

	client = redis.NewClient(&redis.Options{
		// TODO allow custom ports
		// - do the same for postgres too
		Addr:     os.Getenv("REDIS_HOST") + ":6379",
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	return client
}

func Set(key string, value interface{}, expiry time.Duration) error {
	client := getInstance()

	res := client.Set(context.Background(), key, value, expiry)
	return res.Err()
}

func Exists(key string) bool {
	client := getInstance()

	res := client.Exists(context.Background(), key)
	return res.Val() != 0
}
