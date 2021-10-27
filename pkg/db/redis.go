package db

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var REDIS *redis.Client

func InitRedis() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "0.0.0.0:6379"
	}
	REDIS = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := REDIS.Ping().Result()
	if err != nil {
		panic(err)
	}
}
