package redisstore

import (
	"context"
	"errors"
	redis "github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"
)

var rdb *redis.Client

const defaultRedisHost = "redis"
const defaultRedisPort = "6379"
const timeOut = 20

func Connect() error {
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = defaultRedisPort
	}
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = defaultRedisHost
	}
	options := redis.Options{
		Addr: redisHost + ":" + redisPort,
	}
	rdb = redis.NewClient(&options)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeOut)
	defer cancel()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New("Redis returned wrong pong: " + pong)
	}
	log.Println("Successfully connected to Redis.")
	return nil
}
